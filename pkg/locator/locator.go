package locator

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gonum.org/v1/gonum/stat"

	"github.com/golang/geo/r2"
	"github.com/smartwms/locservice/pkg/db"
	"github.com/smartwms/locservice/pkg/mqtt"
)

type Locator struct {
	logger *log.Logger
	db     db.DBAccessor

	measures map[string]measure

	sensors map[string]db.Sensor
	anchors map[string]db.Anchor
	tags    []db.Tag
}

type Location r2.Point

type measure map[int64][]chMeasure

type chMeasure struct {
	v int64
	t int64
}

func New(l *log.Logger, db db.DBAccessor) *Locator {
	return &Locator{
		logger: l,
		db:     db,
	}
}

func (l *Locator) Init() error {
	l.logger.Infoln("Initializing logger")

	var err error
	l.sensors, err = l.db.GetSensors()

	for _, v := range l.sensors {
		l.measures[v.ID] = measure{
			37: []chMeasure{},
			38: []chMeasure{},
			39: []chMeasure{},
		}
	}

	if err != nil {
		l.logger.Errorf("Couldn't load sensors: %+v", err)
		return err
	}

	l.logger.Infoln("Loaded sensors")

	l.anchors, err = l.db.GetAnchors()

	if err != nil {
		l.logger.Errorf("Couldn't load anchors: %+v", err)
		return err
	}

	l.logger.Infoln("Loaded anchors")
	return nil
}

func (l *Locator) Run(meas chan mqtt.RawMeasure) {
	for {
		select {
		case m := <-meas:
			l.logger.Debugf("Got new measure: %+v", m)

			if m.TagID == l.tags[0].ID {
				l.measures[m.SensorID][m.Channel] = append(l.measures[m.SensorID][m.Channel], chMeasure{
					v: -m.RSSI,
					t: m.Timestamp,
				})
			}

			for _, measure := range l.measures {
				measure[37] = cleanOldMeasures(measure[37])
				measure[38] = cleanOldMeasures(measure[38])
				measure[39] = cleanOldMeasures(measure[39])
			}

			means := map[string]float64{}

			for sensor, measure := range l.measures {
				_mean := make([]float64, 3)

				for i, ch := range []int64{37, 38, 39} {
					for _, m := range measure[ch] {
						_mean[i] += float64(m.v)
					}
					_mean[i] /= float64(len(measure[ch]))
				}

				means[sensor] = stat.Mean(_mean, nil)
			}
		}
	}
}

func cleanOldMeasures(target []chMeasure) []chMeasure {
	newMeasures := []chMeasure{}
	threshold := time.Now().UTC().Unix() - 5

	for _, v := range target {
		if v.t > threshold {
			newMeasures = append(newMeasures, v)
		}
	}

	return newMeasures
}

// func (l *Locator) Locate(ld []LocationInputData) Location {
// 	if len(ld) < 3 {
// 		l.logError(fmt.Sprintf("Needs 3 sensors, got %d", len(ld)), nil)
// 		return Location{}
// 	}

// 	sensors, err := loadSensors()

// 	if err != nil {
// 		l.logError("[loadSensors]", err)
// 		return Location{}
// 	}

// 	sA := sensors[ld[0].Sensor]
// 	sB := sensors[ld[1].Sensor]
// 	sC := sensors[ld[2].Sensor]

// 	dists := []float64{
// 		math.Pow(10, (ld[0].RSSI-sA.RefPower)/(-10*sA.AttenuationFactor)),
// 		math.Pow(10, (ld[1].RSSI-sB.RefPower)/(-10*sB.AttenuationFactor)),
// 		math.Pow(10, (ld[2].RSSI-sC.RefPower)/(-10*sC.AttenuationFactor)),
// 	}

// 	m1 := mat.NewDense(2, 2, []float64{
// 		2 * (sA.Position.X - sC.Position.X),
// 		2 * (sA.Position.Y - sC.Position.Y),
// 		2 * (sB.Position.X - sC.Position.X),
// 		2 * (sB.Position.Y - sC.Position.Y),
// 	})

// 	err = m1.Inverse(m1)

// 	if err != nil {
// 		l.logError("[loadSensors]", err)
// 		return Location{}
// 	}

// 	m2 := mat.NewDense(2, 1, []float64{
// 		sA.Position.X*sA.Position.X - sC.Position.X*sC.Position.X + sA.Position.Y*sA.Position.Y - sC.Position.Y*sC.Position.Y + dists[2]*dists[2] - dists[0]*dists[0],
// 		sB.Position.X*sB.Position.X - sC.Position.X*sC.Position.X + sB.Position.Y*sB.Position.Y - sC.Position.Y*sC.Position.Y + dists[2]*dists[2] - dists[1]*dists[1],
// 	})

// 	var pd mat.Dense
// 	pd.Mul(m1, m2)

// 	l.logger.Infof("%+v", pd)

// 	return Location{
// 		X: pd.At(0, 0),
// 		Y: pd.At(1, 0),
// 	}
// }

// func (l *Locator) logError(msg string, err error) {
// 	if l.logger == nil {
// 		return
// 	}

// 	if err == nil {
// 		l.logger.Error(msg)
// 	} else if msg == "" {
// 		l.logger.Error(err)
// 	} else {
// 		l.logger.Errorf("%s: %s", msg, err)
// 	}
// }
