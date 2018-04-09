package main

import (
	"errors"
)

// Item represents an item
type Item struct {
	Name       string
	Action     string
	ItemForUse int
	Contains   []int
}

// FindItemByName finds an item from its name
func FindItemByName(itemName string) (int, *Item, error) {
	for index, item := range Items {
		if item.Name == itemName {
			return index, item, nil
		}
	}
	return -1, nil, errors.New("Item not found")
}

// OpenItem attemps to let a character open an item
func OpenItem(pla *Character, itemName string) {
	loc := Locations[player.CurrentLocation]
	for _, itm := range loc.Items {
		if Items[itm].Name == itemName {
			if Items[itm].ItemForUse != 0 && PlayerHasItem(pla, Items[itm].ItemForUse) {
				loc.Items = append(loc.Items, Items[itm].Contains...)
				Items[itm].Contains = *new([]int)
			} else {
				Output("red", "Could not open the "+itemName)
				return
			}
		} else {
			Output("red", "Could not open the "+itemName)
		}
	}
}

// PlayerHasItem returns whether a character owns a certain item
func PlayerHasItem(pla *Character, itm int) bool {
	for _, pitm := range pla.Items {
		if pitm == itm {
			return true
		}
	}
	return false
}

// RemoveItemFromRoom removes the item from a given location
func (it *Item) RemoveItemFromRoom(loc *Location) {
	for index, itm := range loc.Items {
		if Items[itm].Name == it.Name {
			loc.Items = append(loc.Items[:index], loc.Items[index+1:]...)
		}
	}
}

// ItemInRoom returns whether item is in a given location
//To be refactored on a location struct
func (it *Item) ItemInRoom(loc *Location) bool {
	for _, itm := range loc.Items {
		if Items[itm].Name == it.Name {
			return true
		}
	}
	return false
}

// ItemOnPlayer returns whether the item belongs to a given character
//To be refactored on a character struct
func (it *Item) ItemOnPlayer(pla *Character) bool {
	for _, itm := range pla.Items {
		if Items[itm].Name == it.Name {
			return true
		}
	}
	return false
}

//To be refactored on a location struct
/*
func describeItems(player Character) {
	l := Locations[player.CurrentLocation]

	Output("You see:")
	for _, itm := range l.Items {
		Outputf("\t%s\n", Items[itm].Name)
	}
}
*/
