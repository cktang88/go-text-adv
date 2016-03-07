package location

import "testing"

func TestNewExit(t *testing.T) {
	ex := NewExit(WEST, &Location{})

	if ex == nil {
		t.Errorf("NewExit returned a nil expected an Exit!!")
	}

	if ex.Direction != WEST {
		t.Errorf("NewExit returned wrong Direction, got %s, expected %s", ex.Direction, WEST)
	}

	if directionNames[ex.Direction][0] != "West" {
		t.Errorf("NewExit returned wrong DirectionName, got %s, expected %s", directionNames[ex.Direction][0], "West")
	}

	if directionNames[ex.Direction][1] != "W" {
		t.Errorf("NewExit returned wrong DirectionName, got %s, expected %s", directionNames[ex.Direction][1], "W")
	}

	if ex.LeadsTo.Title != "" {
		t.Errorf("NewExit returned wrong LeadsTo, got %s, expected %s", ex.LeadsTo, nil)
	}

}