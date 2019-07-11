package main

import (
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/bchoates/sensor-simulator/app"
	nats "github.com/nats-io/go-nats"
)

const (
	defaultSensorSubject = "sensorData"
	defaultSensorCount   = 3
	defaultInterval      = 1000
	defaultPrintValues   = true
	defaultNatsAddress   = nats.DefaultURL
	defaultSensorMean    = 50
	defaultSensorDev     = 3
)

func main() {
	sensorSubject := flag.String("sub", defaultSensorSubject, "subject to publish to")
	sensorCount := flag.Uint64("count", defaultSensorCount, "number of sensors to use")
	interval := flag.Int64("interval", defaultInterval, "interval in milliseconds to log values")
	//printValues := flag.Bool("print", defaultPrintValues, "adds a subscriber that prints values added to the bus")
	natsAddress := flag.String("nats", defaultNatsAddress, "nats address")
	sensorMean := flag.Float64("mean", defaultSensorMean, "mean value for the sensors")
	sensorDev := flag.Float64("dev", defaultSensorDev, "standard deviation for the sensors")
	flag.Parse()

	bus, err := nats.Connect(*natsAddress)
	if err != nil {
		log.Fatalf("failed to connect to nats bus: %v", err)
	}

	sim := app.Simulator{}
	for i := uint64(0); i < *sensorCount; i++ {
		name := "sensor" + strconv.FormatUint(i, 10)
		sim.AddSensor(name, *sensorMean, *sensorDev)
	}

	for {
		data, err := sim.Log()
		if err != nil {
			log.Printf("failed to read sensor data: %v", err)
			continue
		}
		err := bus.Publish(*sensorSubject, data)
		if err != nil {
			log.Printf("failed to publish data: %v", err)
			continue
		}
		<-time.After(*interval * time.Millisecond)
	}
}
