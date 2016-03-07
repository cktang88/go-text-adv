package location

import (
	"testing"
)

func TestNewWorld(t *testing.T) {
	loc1 := NewLocation("Entrance", "Entrance of a large cave")
	loc2 := NewLocation("Cave", "inside large cave")
	ex1 := NewExit(NORTH, loc2)
	ex2 := NewExit(SOUTH, loc1)
	loc1.AddExit(ex1)
	loc2.AddExit(ex2)
	theMap := make([]Location, 0)
	theMap = append(theMap, *loc1)
	world, err := NewWorld(theMap)
	if err == nil {
		t.Errorf("NewWorld returned ok when there is not enough rooms to create the world!!")
	}

	theMap = append(theMap, *loc2)
	world, err = NewWorld(theMap)
	if err != nil {
		t.Errorf("NewWorld returned err when it should have created the world")
	}

	if world.Turn != 1 {
		t.Errorf("NewWorld did not set turn == 1")
	}

	//l1 := NewLocation("Bottom of the well", "You have reached the bottom of a deep and rather smelly well. Less than a foot of water remains, and it looks undrinkable.")
	//l2 := NewLocation("Courtyard", "At the centre of the courtyard is an old stone well. A strong and sturdy rope is attached to the well, and descends into the darkness. The only other items of interest is the farmhouse to the north, and a path to the east.")
	//l3 := NewLocation("Farmhouse entrance", "The door to the farmhouse hangs crooked, and is slightly ajar. Obviously no-one has lived here for some time, and you can only guess at what lies within.")
	//l4 := NewLocation("Blood-stained room", "Dried blood stains can be seen on the walls and stone floor of the farmhouse. Whatever massacre occured here long ago, you can only guess. With the abscence of bodies, however, you may never know.")
	//l5 := NewLocation("Long windy path", "You are standing on a long, windy path, leading from the mountains in the far east, to a small farm that lies to the west.")
	//l6 := NewLocation("Base of the mountain", "At the base of the mountain is a path that leads westward beyond a large boulder. Climbing such a mountain would be difficult - if not impossible.")
	//l7 := NewLocation("Top of the mountain", "From this vantage point, you can see all that lies on the plains below. Large boulders dot the landscape, and just within view to the west you make out some sort of a building - though its details are too hard to make out from this distance.")
	//
	//// Create exit objects
	//e1 := NewExit(UP, l2);
	//e2 := NewExit(DOWN, l1);
	//e3 := NewExit(NORTH, l3);
	//e4 := NewExit(SOUTH, l2);
	//e5 := NewExit(NORTH, l4);
	//e6 := NewExit(SOUTH, l3);
	//e7 := NewExit(EAST, l5);
	//e8 := NewExit(WEST, l2);
	//e9 := NewExit(EAST, l6);
	//e10 := NewExit(WEST, l5);
	//e11 := NewExit(UP, l7);
	//e12 := NewExit(DOWN, l6);
	//
	//l1.AddExit(e1);
	//l2.AddExit(e2);
	//l2.AddExit(e3);
	//l2.AddExit(e7);
	//l3.AddExit(e4);
	//l3.AddExit(e5);
	//l4.AddExit(e6);
	//l5.AddExit(e8);
	//l5.AddExit(e9);
	//l6.AddExit(e10);
	//l6.AddExit(e11);
	//l7.AddExit(e12);
	//
	//theMap = make([]Location, 0)
	//theMap = append(theMap, *l1, *l2, *l3, *l4, *l5, *l6, *l7)
	//world, _ = NewWorld(theMap)
	//world.CurrentLocation = *l2
	//world.CurrentLocation.ShowLocation()
	//err = world.Move("D")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	world.CurrentLocation.ShowLocation()
	//}
	//err = world.Move("U")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	world.CurrentLocation.ShowLocation()
	//}
	//err = world.Move("N")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	world.CurrentLocation.ShowLocation()
	//}
}