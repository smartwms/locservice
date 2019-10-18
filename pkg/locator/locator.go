package locator

import (
	"fmt"
	"math"

	log "github.com/sirupsen/logrus"
	"gonum.org/v1/gonum/mat"

	"github.com/golang/geo/r2"
)

type Locator struct {
	logger *log.Logger
}

type LocationInputData struct {
	Sensor string
	RSSI   float64
	AOA    float64
}

type Location r2.Point

func New(l *log.Logger) *Locator {
	return &Locator{
		logger: l,
	}
}

func (l *Locator) Locate(ld []LocationInputData) Location {
	if len(ld) < 3 {
		l.logError(fmt.Sprintf("Needs 3 sensors, got %d", len(ld)), nil)
		return Location{}
	}

	sensors, err := loadSensors()

	if err != nil {
		l.logError("[loadSensors]", err)
		return Location{}
	}

	sA := sensors[ld[0].Sensor]
	sB := sensors[ld[1].Sensor]
	sC := sensors[ld[2].Sensor]

	dists := []float64{
		math.Pow(10, (ld[0].RSSI-sA.RefPower)/(-10*sA.AttenuationFactor)),
		math.Pow(10, (ld[1].RSSI-sB.RefPower)/(-10*sB.AttenuationFactor)),
		math.Pow(10, (ld[2].RSSI-sC.RefPower)/(-10*sC.AttenuationFactor)),
	}

	m1 := mat.NewDense(2, 2, []float64{
		2 * (sA.Position.X - sC.Position.X),
		2 * (sA.Position.Y - sC.Position.Y),
		2 * (sB.Position.X - sC.Position.X),
		2 * (sB.Position.Y - sC.Position.Y),
	})

	err = m1.Inverse(m1)

	if err != nil {
		l.logError("[loadSensors]", err)
		return Location{}
	}

	m2 := mat.NewDense(2, 1, []float64{
		sA.Position.X*sA.Position.X - sC.Position.X*sC.Position.X + sA.Position.Y*sA.Position.Y - sC.Position.Y*sC.Position.Y + dists[2]*dists[2] - dists[0]*dists[0],
		sB.Position.X*sB.Position.X - sC.Position.X*sC.Position.X + sB.Position.Y*sB.Position.Y - sC.Position.Y*sC.Position.Y + dists[2]*dists[2] - dists[1]*dists[1],
	})

	var pd mat.Dense
	pd.Mul(m1, m2)

	l.logger.Infof("%+v", pd)

	return Location{
		X: pd.At(0, 0),
		Y: pd.At(1, 0),
	}
}

func (l *Locator) logError(msg string, err error) {
	if l.logger == nil {
		return
	}

	if err == nil {
		l.logger.Error(msg)
	} else if msg == "" {
		l.logger.Error(err)
	} else {
		l.logger.Errorf("%s: %s", msg, err)
	}
}
