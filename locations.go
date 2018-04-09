package main

import (
	"errors"
	"strings"
)

// Location represents a location
type Location struct {
	Description string
	Transitions []string
	Events      []string
	Items       []int
}

// CanGoTo returns whether the location can be visited
func (loc *Location) CanGoTo(locName string) bool {
	for _, name := range loc.Transitions {
		if (strings.ToLower(name) == locName) || (strings.ToLower(name[0:3]) == locName[0:3]) {
			return true
		}
	}
	return false
}

// FindLocationName returns a location from its name
func FindLocationName(inputName string) (string, error) {
	for key, _ := range Locations {
		if (strings.ToLower(key) == inputName) || (strings.ToLower(key[0:3]) == inputName[0:3]) {
			return key, nil
		}
	}
	return "", errors.New("Can't find localtion " + inputName)
}
