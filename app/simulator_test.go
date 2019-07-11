package app

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestAddSensor(t *testing.T) {
	name := "testsensor"
	mean := 50.0
	dev := 5.0
	ts := &sensor{
		Name:   name,
		Mean:   mean,
		StdDev: dev,
		seed:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	sim := &Simulator{}
	sim.AddSensor(name, mean, dev)

	if got, want := ts, sim.sensors[0]; !reflect.DeepEqual(got, want) {
		t.Errorf("structs differ. got %v, want %v", got, want)
	}
}

func TestAddSensorMultiple(t *testing.T) {
	tt := []struct {
		Name string
		Mean float64
		Dev  float64
	}{
		{"sensor1", 50.0, 5.0},
		{"sensor2", 50.0, 5.0},
		{"sensor3", 50.0, 5.0},
	}
	sim := &Simulator{}
	for i, test := range tt {
		ts := &sensor{
			Name:   test.Name,
			Mean:   test.Mean,
			StdDev: test.Dev,
			seed:   rand.New(rand.NewSource(time.Now().UnixNano())),
		}
		sim.AddSensor(test.Name, test.Mean, test.Dev)
		if got, want := ts, sim.sensors[i]; !reflect.DeepEqual(got, want) {
			t.Errorf("structs differ. got %v, want %v", got, want)
		}
	}
}
