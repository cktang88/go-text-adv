package location

import (
	"fmt"
)

type Location struct {
	Title 		string
	Description 	string
	Exits 		[]*Exit
}

func NewLocation(title string, desc string) *Location {
	loc := new(Location)
	loc.Title = title
	loc.Description = desc
	loc.Exits = make([]*Exit, 0)
	return loc
}

func (loc *Location) AddExit( exit *Exit ) {
	loc.Exits = append(loc.Exits, exit)
}

func (loc *Location) GetExits() ([]*Exit) {
	return loc.Exits
}

func (loc *Location) ShortDirections() []string {
	dir := make([]string, 0)
	for _, ex := range loc.Exits {
		dir = append(dir, ex.GetShortName())
	}
	return dir
}

func (loc *Location) ShowLocation() {
	fmt.Println(loc.Title)
	fmt.Println(loc.Description)
	//fmt.Println("Exits:")
	//for _, ex := range loc.Exits {
	//	fmt.Print("\t" + ex.ToString())
	//
	//	if(ex.LeadsTo.Title == "") {
	//		fmt.Println(": leads nowhere!!!")
	//	} else {
	//		fmt.Println(": leads to " + ex.LeadsTo.Title)
	//	}
	//}
}