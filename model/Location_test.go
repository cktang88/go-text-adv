package location

import "testing"

func TestNewLocation(t *testing.T) {
	loc := NewLocation("Room5", "description")
	ex := NewExit( WEST, loc)
	loc.AddExit(ex)

	if loc == nil {
		t.Errorf("NewExit returned a nil expected an Exit!!")
	}

	exits := loc.GetExits()
	if len(exits) != 1 {
		t.Errorf("Not the correct amount of exits in this location, Got %s Expected %s\n", len(exits), 1)
	}
	//loc.ShowLocation()
}