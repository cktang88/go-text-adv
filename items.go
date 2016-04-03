package main

import "fmt"

type Item struct {
	Name string
	ItemForUse int
	Contains []int
}

var Items = map[int]*Item{
	1: { Name: "Master Key" },
	2: { Name: "Chest", ItemForUse: 1, Contains: []int{3}},
	3: { Name: "Medal" },
}

var locMap = map[string]*Location{
	"main": { Description: "You are on the bridge of a spaceship sitting in the Captain's chair.", Items: []int{1, 2} },
}

func ProcessInput(input int) {
	itm := Items[input]
	fmt.Printf("%d - %v\n", input, itm)
	if Items[input].ItemForUse != 0 {
		DisplayInfo("See if user has item to use..")
	}
}

func describeItems(loc string) {
	l := locMap["main"]
	input := 0
	for input != 3 {
		DisplayInfo("You see:")
		for _, itm := range l.Items {
			DisplayInfof("\t%s\n", Items[itm].Name)
		}

		DisplayInfo("1. Pickup Key")
		DisplayInfo("2. Open Chest")
		DisplayInfo("3. Quit")
		GetUserInput(&input)
		ProcessInput(input)
	}
}

var player Character

//func main() {
//	player = *new(Character)
//
//	describeItems("main")
//}

