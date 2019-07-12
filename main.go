package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/bchoates/sensor-simulator/app"
	nats "github.com/nats-io/nats.go"
)

const (
	defaultSensorSubject = "sensorData"
	defaultSensorCount   = 3
	defaultInterval      = 1000
	defaultPrintValues   = false
	defaultNatsAddress   = nats.DefaultURL
	defaultSensorMean    = 50
	defaultSensorDev     = 3
)

type event struct {
	Name      string  `json:"name"`
	TimeStamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

func main() {
	sensorSubject := flag.String("sub", defaultSensorSubject, "subject to publish to")
	sensorCount := flag.Int64("count", defaultSensorCount, "number of sensors to use")
	interval := flag.Int64("interval", defaultInterval, "interval in milliseconds to log values")
	printValues := flag.Bool("print", defaultPrintValues, "adds a subscriber that prints values added to the bus")
	natsAddress := flag.String("nats", defaultNatsAddress, "nats address")
	sensorMean := flag.Float64("mean", defaultSensorMean, "mean value for the sensors")
	sensorDev := flag.Float64("dev", defaultSensorDev, "standard deviation for the sensors")
	flag.Parse()

	bus, err := nats.Connect(*natsAddress)
	if err != nil {
		log.Fatalf("failed to connect to nats bus: %v", err)
	}

	sim := app.Simulator{}
	events := make([]event, *sensorCount)

	for i := int64(0); i < *sensorCount; i++ {
		name := "sensor" + strconv.FormatInt(i, 10)
		sim.AddSensor(name, *sensorMean, *sensorDev, i)
		events[i].Name = name
	}

	if *printValues {
		go func(conn *nats.Conn, subject string) {
			conn.Subscribe(subject, func(m *nats.Msg) {
				fmt.Printf("sensor data rec: %v\n", string(m.Data))
			})
		}(bus, *sensorSubject)
	}

	for {
		timestamp := time.Now().Unix()
		data, err := sim.Log()
		if err != nil {
			log.Printf("failed to read sensor data: %v", err)
			continue
		}
		for i, val := range data {
			events[i].TimeStamp = timestamp
			events[i].Value = (math.Round(val*100) / 100)
		}

		message, err := json.Marshal(events)
		if err != nil {
			log.Printf("error marashalling data")
		}

		err = bus.Publish(*sensorSubject, message)
		if err != nil {
			log.Printf("failed to publish data: %v", err)
			continue
		}
		<-time.After(time.Duration(*interval) * time.Millisecond)
	}
}
