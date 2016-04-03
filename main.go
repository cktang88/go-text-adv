package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
)

var Out *os.File
var In *os.File


func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	Out = os.Stdout
	In = os.Stdin
}

func main() {
	//Player
	player := new(Character)
	player.Name = "Paul"
	player.Speed = 1 + rand.Intn(100)
	player.Health = 100
	player.Alive = true
	player.Weap = 1
	player.CurrentLocation = "Bridge"

	player.Play()
}

func DisplayInfof(format string, args ...interface{}) {
	fmt.Fprintf(Out, format, args...)
}

func DisplayInfo(args ...interface{}) {
	fmt.Fprintln(Out, args...)
}

func GetUserInput(i *int) {
	fmt.Fscan(In, i)
}

