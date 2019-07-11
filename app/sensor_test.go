package app

import (
	"math/rand"
	"testing"
)

func TestValue(t *testing.T) {
	s := sensor{
		Name: "testsensor",
		seed: rand.New(rand.NewSource(1)),
	}

	if _, err := s.Value(); err != nil {
		t.Errorf("error getting value")
	}
}

func TestValueNoSeed(t *testing.T) {
	s := sensor{
		Name: "testsensor",
	}

	if _, err := s.Value(); err == nil {
		t.Errorf("error getting value")
	}
}
