package main

import (
	"math/rand"
	"time"
)


type Event struct {
	Type        string
	Chance      int
	Description string
	Health      int
	Evt         string
}


var evts = map[string]*Event{
	"alienAttack":     {Type: "Combat", Chance: 20, Description: "An alien beams in front of you and shoots you with a ray gun.", Health: -50, Evt: "doctorTreatment"},
	"doctorTreatment": {Type: "Story", Chance: 10, Description: "The doctor rushes in and inject you with a health boost.", Health: +30, Evt: ""},
	"android":         {Type: "Story", Chance: 50, Description: "Data is in the turbo lift and says hi to you", Health: 0, Evt: ""},
	"relaxing":        {Type: "Story", Chance: 100, Description: "In the lounge you are so relaxed that your health improves.", Health: +10, Evt: ""},
}

func (e *Event) ProcessEvent(player *Character) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	if e.Chance >= r1.Intn(100) {
		if e.Type == "Combat" {
			//Generate opponent

			opp := new(Character)
			*opp = *Enemies[1+rand.Intn(len(Enemies)-1)]
			opp.Npc = true
			opp.Speed = 1 + rand.Intn(100)
			DisplayInfo("A " + opp.Name + " jumps in front of you and attacks")

			players := Players{*player, *opp}
			runBattle(players)
			// Run combat
			DisplayInfo("Combat Event")
		} else {
			DisplayInfo("\t" + e.Description)
			if e.Evt != "" {
				e.Health = e.Health + evts[e.Evt].ProcessEvent(player)
			}
		}
		return e.Health
	}
	return 0
}

