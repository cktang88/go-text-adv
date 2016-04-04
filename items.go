package main

import (
	"strings"
)

type Item struct {
	Name string
	Action string
	ItemForUse int
	Contains []int
}

var Items = map[int]*Item{
	1: { Name: "Key" },
	2: { Name: "Chest", ItemForUse: 1, Contains: []int{3}},
	3: { Name: "Medal" },
}

var locMap = map[string]*Location{
	"main": { Description: "You are on the bridge of a spaceship sitting in the Captain's chair.", Items: []int{1, 2} },
}

func ProcessCommands(player Character, input string) {
	tokens := strings.Fields(input)
	command := tokens[0]
	itemName := ""
	if len(tokens) > 1 {
		itemName = tokens[1]
	}
	DisplayInfo(tokens)
	loc := locMap[player.CurrentLocation]
	switch(command) {
	case "get":
		//Make sure we do not pick it up twice
		if ItemInRoom(loc, itemName)  && !ItemOnPlayer(player, itemName) {
			PutItemInPlayerBag(player, itemName)
			ItemRemoveFromRoom(loc, itemName)
		} else {
			DisplayInfo("Could not get " + itemName)
		}
	case "open":
		OpenItem(player, itemName)
	case "inv":
		DisplayInfo("Your Inventory: ")
		for _, itm := range player.Items {
			DisplayInfo("\t" + Items[itm].Name)
		}
	default:
	}
}

func OpenItem(pla Character, itemName string) {
	DisplayInfo("Opening " + itemName)
	loc := locMap[player.CurrentLocation]
	for _, itm := range loc.Items {
		if Items[itm].Name == itemName {
			DisplayInfo("Loop >> " + Items[itm].Name)
			if Items[itm].ItemForUse != 0 && PlayerHasItem(pla, Items[itm].ItemForUse) {
				loc.Items = append(loc.Items, Items[itm].Contains...)
				Items[itm].Contains = *new([]int)
			}
		} else {
			DisplayInfo("Could not open the " + itemName)
		}
	}
}

func PlayerHasItem(pla Character, itm int) bool {
	for _, pitm := range pla.Items {
		if pitm == itm {
			return true
		}
	}
	return false
}
//To be refactored on a character struct
func PutItemInPlayerBag(pla Character, itemName string) {
	for index, itm := range Items {
		if itm.Name == itemName {
			player.Items = append(player.Items, index)
			return
		}
	}
}


//To be refactored on a location struct
func ItemRemoveFromRoom(loc *Location, itemName string) {
	for index, itm := range loc.Items {
		if Items[itm].Name == itemName{
			loc.Items = append(loc.Items[:index], loc.Items[index+1:]...)
		}
	}
}

//To be refactored on a location struct
func ItemInRoom(loc *Location, itemName string) bool {
	for _, itm := range loc.Items {
		if Items[itm].Name == itemName{
			return true
		}
	}
	return false
}

//To be refactored on a character struct
func ItemOnPlayer(pla Character, itemName string) bool {
	for _, itm := range pla.Items {
		if Items[itm].Name == itemName{
			return true
		}
	}
	return false
}

//To be refactored on a location struct
func describeItems(player Character) {
	l := locMap[player.CurrentLocation]

	DisplayInfo("You see:")
	for _, itm := range l.Items {
		DisplayInfof("\t%s\n", Items[itm].Name)
	}
}

var player Character

func main() {
	player = *new(Character)
	player.CurrentLocation = "main"
	input := ""
	for input != "quit" {
		describeItems(player)
		input = GetUserStrInput()
		ProcessCommands(player, input)
	}
}
