package app

import (
	"math/rand"
	"time"
)

//Simulator structure that holds sensors and provides method to output sensor data
type Simulator struct {
	sensors []*sensor
}

//AddSensor adds sensor to the list
func (sim *Simulator) AddSensor(name string, mean, dev float64, random int64) {
	sen := &sensor{
		Name:   name,
		Mean:   mean,
		StdDev: dev,
		seed:   rand.New(rand.NewSource(time.Now().UnixNano() + random)),
	}
	sim.sensors = append(sim.sensors, sen)
}

//Log retrieves data from the sensors
func (sim *Simulator) Log() ([]float64, error) {
	values := make([]float64, len(sim.sensors))
	for i, s := range sim.sensors {
		val, err := s.Value()
		if err != nil {
			return nil, err
		}
		values[i] = val
	}
	return values, nil
}
