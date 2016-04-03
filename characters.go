package main


type Character struct {
	Name    string
	Health  int
	Evasion int
	Alive   bool
	Speed   int
	Weap    int
	Npc     bool
	Itens 	[]Item

	Welcome         string
	CurrentLocation string
}

var Enemies = map[int]*Character{
	0: {Name: "Klingon", Health: 50, Alive: true, Weap: 2, Npc: true},
	1: {Name: "Romulan", Health: 55, Alive: true, Weap: 3, Npc: true},
}

func (p *Character) Equip(w int) {
	p.Weap = w
}

func (p *Character) Attack() int {
	return Weaps[p.Weap].Fire()
}

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


func (p *Character) Play() {
	DisplayInfo(p.Welcome)
	for {
		DisplayInfo(LocationMap[p.CurrentLocation].Description)
		p.ProcessEvents(LocationMap[p.CurrentLocation].Events)
		if p.Health <= 0 {
			DisplayInfo("You are dead, game over!!!")
			return
		}
		DisplayInfo("Health:", p.Health)
		DisplayInfo("You can go to these places:")
		for index, loc := range LocationMap[p.CurrentLocation].Transitions {
			DisplayInfof("%d - %s\n", index+1, loc)
		}
		i := 0
		for i < 1 || i > len(LocationMap[p.CurrentLocation].Transitions) {
			DisplayInfof("%s%d%s\n", "Where do you want to go (0 - to quit), [1...", len(LocationMap[p.CurrentLocation].Transitions), "]: ")
			GetUserInput(&i)
		}
		newLoc := i - 1
		p.CurrentLocation = LocationMap[p.CurrentLocation].Transitions[newLoc]

	}
}

func (p *Character) ProcessEvents(events []string) {
	for _, evtName := range events {
		p.Health += evts[evtName].ProcessEvent(p)
	}
}

