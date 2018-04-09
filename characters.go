package main

// Character represents a character
type Character struct {
	Name    string
	Health  int
	Evasion int
	Alive   bool
	Speed   int
	Weap    int
	Npc     bool
	Items   []int

	Welcome         string
	CurrentLocation string
}

// Equip equips an active weapon to the character
func (p *Character) Equip(w int) {
	p.Weap = w
}

// Attack fires all the player's weapons
func (p *Character) Attack() int {
	return Weaps[p.Weap].Fire()
}

// Players is an array of Characters
type Players []Character

func (slice Players) Len() int {
	return len(slice)
}

func (slice Players) Less(i, j int) bool {
	return slice[i].Speed > slice[j].Speed //Sort descending
	//return slice[i].Speed < slice[j].Speed;		//Sort ascending
}

func (slice Players) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Play acts out one player turn
func (p *Character) Play() {
	Output(p.Welcome)
	for {
		Output("blue", Locations[p.CurrentLocation].Description)
		p.ProcessEvents(Locations[p.CurrentLocation].Events)
		if p.Health <= 0 {
			Output("white", "You are dead, game over!!!")
			return
		}
		Output("blue", "Health:", p.Health)
		if len(Locations[p.CurrentLocation].Items) > 0 {
			Output("yellow", "You can see:")
			for _, itm := range Locations[p.CurrentLocation].Items {
				Outputf("yellow", "\t%s", Items[itm].Name)
			}
		}
		Output("green", "You can go to these places:")
		for _, loc := range Locations[p.CurrentLocation].Transitions {
			Outputf("green", "\t%s", loc)
		}
		cmd := UserInputln()
		ProcessCommands(p, cmd)
	}
}

func (p *Character) ProcessEvents(events []string) {
	for _, evtName := range events {
		p.Health += Events[evtName].ProcessEvent(p)
	}
}
