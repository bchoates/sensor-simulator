package app

import (
	"fmt"
	"math/rand"
)

type sensor struct {
	Name   string
	Mean   float64
	StdDev float64
	seed   *rand.Rand
}

func (s *sensor) Value() (float64, error) {
	if s.seed == nil {
		return 0, fmt.Errorf("seed is nil")
	}
	return (s.seed.NormFloat64()*s.StdDev + s.Mean), nil
}
