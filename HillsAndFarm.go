package main

import (
	"strings"
	"os"
	"bufio"

	. "gotextadv/model"
	"fmt"
)

// https://play.golang.org/p/1UhR5V9_Et
// http://codereview.stackexchange.com/questions/36768/tiny-text-adventure

var reader *bufio.Reader

func init() {
	reader = bufio.NewReader(os.Stdin)
}

func main() {
	l1 := NewLocation("Bottom of the well", "You have reached the bottom of a deep and rather smelly well. Less than a foot of water remains, and it looks undrinkable.")
	l2 := NewLocation("Courtyard", "At the centre of the courtyard is an old stone well. A strong and sturdy rope is attached to the well, and descends into the darkness. The only other items of interest is the farmhouse to the north, and a path to the east.")
	l3 := NewLocation("Farmhouse entrance", "The door to the farmhouse hangs crooked, and is slightly ajar. Obviously no-one has lived here for some time, and you can only guess at what lies within.")
	l4 := NewLocation("Blood-stained room", "Dried blood stains can be seen on the walls and stone floor of the farmhouse. Whatever massacre occured here long ago, you can only guess. With the abscence of bodies, however, you may never know.")
	l5 := NewLocation("Long windy path", "You are standing on a long, windy path, leading from the mountains in the far east, to a small farm that lies to the west.")
	l6 := NewLocation("Base of the mountain", "At the base of the mountain is a path that leads westward beyond a large boulder. Climbing such a mountain would be difficult - if not impossible.")
	l7 := NewLocation("Top of the mountain", "From this vantage point, you can see all that lies on the plains below. Large boulders dot the landscape, and just within view to the west you make out some sort of a building - though its details are too hard to make out from this distance.")

	// Create exit objects
	e1 := NewExit(UP, l2);
	e2 := NewExit(DOWN, l1);
	e3 := NewExit(NORTH, l3);
	e4 := NewExit(SOUTH, l2);
	e5 := NewExit(NORTH, l4);
	e6 := NewExit(SOUTH, l3);
	e7 := NewExit(EAST, l5);
	e8 := NewExit(WEST, l2);
	e9 := NewExit(EAST, l6);
	e10 := NewExit(WEST, l5);
	e11 := NewExit(UP, l7);
	e12 := NewExit(DOWN, l6);

	l1.AddExit(e1);
	l2.AddExit(e2);
	l2.AddExit(e3);
	l2.AddExit(e7);
	l3.AddExit(e4);
	l3.AddExit(e5);
	l4.AddExit(e6);
	l5.AddExit(e8);
	l5.AddExit(e9);
	l6.AddExit(e10);
	l6.AddExit(e11);
	l7.AddExit(e12);

	theMap := make([]Location, 0)
	theMap = append(theMap, *l1, *l2, *l3, *l4, *l5, *l6, *l7)

	world, _ := NewWorld(theMap)
	world.CurrentLocation = *l2

	endGame := false
	for !endGame {
		world.CurrentLocation.ShowLocation()
		moveDir := userInput(world.CurrentLocation.ShortDirections())
		if moveDir == "Q" {
			endGame = true
		} else {
			world.Move(moveDir)
		}
	}
	fmt.Println("Game is over thanks for playing...")
}


func userInput(dirs []string) string {
	list := []string{ "N", "S", "W", "E", "U", "D", "I", "O", "Q" }
	valid := false
	for !valid {
		fmt.Printf("What would you like to do %v? ", dirs)
		text, _ := reader.ReadString('\n')
		command := strings.Split(text, " ")
		command[0] = strings.ToUpper(strings.TrimSpace(command[0]))
		if stringInSlice(command[0], list) {
			return command[0]
		}
	}
	return ""
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}