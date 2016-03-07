package main

import (
	"fmt"
	"strings"
	"time"
	"bufio"
	"os"
)

const (
	NORTH = 0
	EAST = 1
	SOUTH = 2
	WEST = 3
)
// http://codereview.stackexchange.com/questions/36768/tiny-text-adventure

type Weapon struct {
	name string
	minAttack int
	maxAttack int
}

type Player struct {
	name 		string
	move 		int
	alive 		bool
	health 		int
	weapon 		Weapon
	location	*Location
}

type Npc struct {
	name 		string
	move 		int
	foe		bool
	minAttack 	int
	maxAttack 	int
}

type Location struct {
	id		int
	name 		string
	desc 		string
	question	string
	responses	[]string
	autoMoveId	*Location
	weapons 	[]Weapon
	npcs		[]Npc
	exits		[4]*Location
}

func (l *Location) AutoMove(p Player) {
	if l.autoMoveId != nil

}

type Event struct {
	quest	string
}
var rooms []Location

type Game struct {
	player 		Player
	turn 		int
}

var game Game
var reader *bufio.Reader

func init() {
	rooms = make([]Location, 0)
	npc1 := Npc{name: "spider", minAttack: 2, maxAttack: 5, foe: true, move: 1 }
	npcs := make([]Npc, 0)
	npcs = append(npcs, npc1)

	room1 := Location{
		id: 1,
		name:"Area 2",
		desc: "As you proceed further into the cave, you see a small glowing object",
		question: "Do you approach the object? [y/n] ",
		npcs: npcs,
	}
	stik := Weapon{ name: "stick", minAttack: 3, maxAttack: 10 }
	weaps := make([]Weapon, 0)
	weaps = append(weaps, stik)

	room0 := Location{
		id: 0,
		name:"Area 1",
		desc: "You enter a dark cavern out of curiosity. It is very dark but the cavern seems to strech out to the North.   You can make out a small stick on the floor.",
		question: "Do you pick up the stick? [y/n]: ",
		weapons: weaps,
	}
	room0.responses = []string{ "Y", "YES" }
	rooms = append(rooms, room0)
	rooms = append(rooms, room1)

	rooms[0].exits[NORTH] = &rooms[1]
	rooms[0].autoMoveId = &rooms[1]
	rooms[1].exits[SOUTH] = &rooms[0]

	game = Game{turn: 1}
	game.player = Player{name: "John", health: 100, location: &rooms[0]}

	reader = bufio.NewReader(os.Stdin)
}

func main() {
	game.player.alive = true
	for game.player.alive {
		if gameLoop() {
			if !userInput2("You managed to escape the cavern alive! Would you like to play again? [y/n]: ", rooms[0].responses) {
				game.player.alive = false
			}
		} else {
			if !userInput2("You have died do you want to play again? [[y/n]?", rooms[0].responses) {
				game.player.alive = false
			}
		}

	}
}

func gameLoop() bool {
	fmt.Println(game.player.name, game.player.location.desc)
	if userInput2(game.player.location.question, game.player.location.responses) {
		fmt.Println("You picked up the stick.")
	} else {
		fmt.Println("You wonder if not grabbing the stick is a mistake but...")
	}

	return false
}

func userInput2(question string, list []string) bool {
	fmt.Print(question)
	time.Sleep(time.Second * 3)
	text, _ := reader.ReadString('\n')
	text = strings.ToUpper(strings.TrimSpace(text))
	if stringInSlice(text, list) {
		return true
	} else { return false }
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
/*
func runGame() bool {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Welcome to the cavern of secrets!")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	time.Sleep(time.Second * 3)

	stick := 0
	if askUser("You enter a dark cavern out of curiosity. It is dark and you can only make out a small stick on the floor.",
		"Do you take it? [y/n]: ") {
		stick = 1
	}
	time.Sleep(time.Second * 2)
	if stick == 1 {
		stick = 1
	}
	if askUser("As you proceed further into the cave, you see a small glowing object", "Do you approach the object? [y/n]") {
		fmt.Println("You approach the object...")
		time.Sleep(time.Second * 2)
		fmt.Println("As you draw closer, you begin to make out the object as an eye!")
		time.Sleep(time.Second * 1)
		if askUser("The eye belongs to a giant spider!", "Do you try to fight it? [Y/N]") {

		} else {

		}

	} else {
		fmt.Println("You turn away from the glowing object, and attempt to leave the cave...")
		time.Sleep(time.Second * 1)
		fmt.Println("But something won't let you....")
		time.Sleep(time.Second * 2)
	}
	return false
}

func askUser(question string, list []string) bool {
	fmt.Println(desc)
	fmt.Print(question)
	text, _ := reader.ReadString('\n')
	text = strings.ToUpper(strings.TrimSpace(text))
	if stringInSlice(text, list) {
		return true
	} else { return false }
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
*/