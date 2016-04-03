package main

type Location struct {
	Description string
	Transitions []string
	Events      []string
	Items 	    []int
}

var LocationMap = map[string]*Location{
	"Bridge":      { Description: "You are on the bridge of a spaceship sitting in the Captain's chair.", Transitions: []string{"Ready Room", "Turbo Lift"}, Events: []string{"alienAttack"}},
	"Ready Room":  { Description: "The Captain's ready room.", Transitions: []string{"Bridge"}, Events: []string{}, Items: []int{2}},
	"Turbo Lift":  { Description: "A Turbo Lift that takes you anywhere in the ship.", Transitions: []string{"Bridge", "Lounge", "Engineering"}, Events: []string{"android"}},
	"Engineering": { Description: "You are in engineering where you see the star drive", Transitions: []string{"Turbo Lift"}, Events: []string{"alienAttack"}, Items: []int{1}},
	"Lounge":      { Description: "You are in the lounge, you feel very relaxed", Transitions: []string{"Turbo Lift"}, Events: []string{"relaxing"}},
}

