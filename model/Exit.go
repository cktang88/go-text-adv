package location

// https://play.golang.org/p/I2qW3nm6-j

const (
	UNDEFINED = 0
	NORTH = 1
	SOUTH = 2
	EAST = 3
	WEST = 4
	UP = 5
	DOWN = 6
	NORTHEAST = 7
	NORTHWEST = 8
	SOUTHEAST = 9
	SOUTHWEST = 10
	IN = 11
	OUT = 12
)

var directionNames = map[int][2]string {
	UNDEFINED: {"UNDEFINED", "null"},
	NORTH: {"North", "N"},
	SOUTH: {"South", "S"},
	EAST: {"East", "E"},
	WEST: {"West", "W"},
	UP: {"Up", "U"},
	DOWN: {"Down", "D"},
	NORTHEAST: {"NorthEast", "NE"},
	NORTHWEST: {"NorthWest", "NW"},
	SOUTHEAST: {"SouthEast", "SE"},
	SOUTHWEST: {"SouthWest", "SW"},
	IN: {"In", "I"},
	OUT: {"Out", "O"},
}


type Exit struct {
	LeadsTo			*Location
	Direction 		int
}


func NewExit(dir int, leadsTo *Location) *Exit {
	ex := new(Exit)
	ex.Direction = dir
	ex.LeadsTo = leadsTo
	return ex
}


func (ex *Exit) GetLongName() string {
	return directionNames[ex.Direction][0]
}


func (ex *Exit) GetShortName() string {
	return directionNames[ex.Direction][1]
}

func (ex *Exit) ToString() string {
	return ex.GetLongName() + "(" + ex.GetShortName() + ")"
}