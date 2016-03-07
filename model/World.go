package location

import (
	"errors"
)

type World struct {
	Turn            int
	CurrentLocation Location
	WorldMap        []Location
}


func NewWorld(wm []Location) (World, error) {
	w := World{Turn: 1}
	if len(wm) > 1 {
		w.WorldMap = wm
		w.CurrentLocation = wm[0]
	} else {
		return w, errors.New("Map does not have enought locations to create a world")
	}
	return w, nil
}


func (w *World)Move(dir string) error {
	for _ , ex := range w.CurrentLocation.Exits {
		if ex.GetShortName() == dir {
			w.CurrentLocation = *ex.LeadsTo
			return nil
		}
	}
	return errors.New("Can't go in that direction")
}