package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/golang/geo/r2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type DB struct {
	c   *firestore.Client
	ctx context.Context
}

type DBAccessor interface {
	GetSensors() (map[string]Sensor, error)
	GetAnchors() (map[string]Anchor, error)
}

type Sensor struct {
	ID       string   `json:"id"`
	Position r2.Point `json:"pos"`
}

type Anchor struct {
	ID       string   `json:"id"`
	Position r2.Point `json:"pos"`
}

type Tag struct {
	ID string `json:"id"`
}

type Measure struct {
	RSSI struct {
		Ch37 int64 `json:"ch37"`
		Ch38 int64 `json:"ch38"`
		Ch39 int64 `json:"ch39"`
	} `json:"rssi"`
	Sensor    string    `json:"sensor"`
	Tag       string    `json:"tag"`
	Timestamp time.Time `json:"timestamp"`
}

type Raw struct {
	SensorID  string `firestore:"sensor"`
	TagID     string `firestore:"tag"`
	Channel   int64  `firestore:"ch"`
	RSSI      int64  `firestore:"rssi"`
	Timestamp int64  `firestore:"ts"`
}

func NewClient() *DB {
	ctx := context.Background()
	sa := option.WithCredentialsFile("keys/smartwms.json")
	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		c:   client,
		ctx: ctx,
	}
}

func (db *DB) GetSensors() (map[string]Sensor, error) {
	iter := db.c.Collection("sensors").Documents(db.ctx)
	defer iter.Stop()

	sensors := map[string]Sensor{}

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
			return nil, err
		}

		posX, posY, _ := posDataToFloats(doc.Data())

		sensors[doc.Ref.ID] = Sensor{
			ID:       doc.Ref.ID,
			Position: r2.Point{X: posX, Y: posY},
		}
	}

	return sensors, nil
}

func (db *DB) GetAnchors() (map[string]Anchor, error) {
	iter := db.c.Collection("anchors").Documents(db.ctx)
	defer iter.Stop()

	anchors := map[string]Anchor{}

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
			return nil, err
		}

		posX, posY, _ := posDataToFloats(doc.Data())

		anchors[doc.Ref.ID] = Anchor{
			ID:       doc.Ref.ID,
			Position: r2.Point{X: posX, Y: posY},
		}
	}

	return anchors, nil
}

func (db *DB) GetTags() ([]Tag, error) {
	iter := db.c.Collection("tags").Documents(db.ctx)
	defer iter.Stop()

	tags := []Tag{}

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
			return nil, err
		}

		tags = append(tags, Tag{
			ID: doc.Ref.ID,
		})
	}

	return tags, nil
}

func (db *DB) AddRawMeasure(sensor, tag string, ts, ch, rssi int64) error {
	_, _, err := db.c.Collection("raw").Add(db.ctx, Raw{
		SensorID:  sensor,
		TagID:     tag,
		Timestamp: ts,
		Channel:   ch,
		RSSI:      rssi,
	})

	return err
}

func (db *DB) GetTagLastMeasures(tag string) ([]Measure, error) {
	iter := db.c.Collection("measures").
		Where("tag", "==", tag).
		Documents(db.ctx)

	defer iter.Stop()

	measures := []Measure{}

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
			return nil, err
		}

		fmt.Printf("%+v\n", doc.Data())

		tmp, _ := json.Marshal(doc.Data())
		meas := Measure{}

		json.Unmarshal(tmp, &meas)
		meas.Sensor = doc.Data()["sensor"].(*firestore.DocumentRef).ID

		measures = append(measures, meas)
	}

	return measures, nil
}

func posDataToFloats(data map[string]interface{}) (float64, float64, error) {
	pos, ok := data["pos"].(map[string]interface{})

	if !ok {
		return 0, 0, errors.New("Couldn't find pos in returned data")
	}

	posX, ok := pos["x"].(float64)

	if !ok {
		posX = float64(pos["x"].(int64))
	}

	posY, ok := pos["x"].(float64)

	if !ok {
		posY = float64(pos["y"].(int64))
	}

	return posX, posY, nil
}
