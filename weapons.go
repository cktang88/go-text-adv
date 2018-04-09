package main

import (
	"math/rand"
	"time"
)

// Weapon is a struct representing a weapon, with min and max attack, and name
type Weapon struct {
	minAtt int
	maxAtt int
	Name   string
}

// Fire fires a Weapon and returns damage dealt
func (w *Weapon) Fire() int {
	return w.minAtt + rand.Intn(w.maxAtt-w.minAtt)
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
