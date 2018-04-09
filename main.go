package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
)

// Out is a pointer to the output file
var Out *os.File

// In is a pointer to the input file
var In *os.File

var player Character

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	Out = os.Stdout
	In = os.Stdin
}

func main() {
	player = *new(Character)
	player.Name = "Paul"
	player.Speed = 1 + rand.Intn(100)
	player.Health = 100
	player.Alive = true
	player.Weap = 1
	player.CurrentLocation = "Bridge"

	player.Play()
}

// Outputf prints args in a given format, with colors
func Outputf(c string, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	Output(c, s)
}

// Output prints args in color
func Output(c string, args ...interface{}) {
	s := fmt.Sprint(args...)
	col := color.WhiteString
	switch c {
	case "green":
		col = color.GreenString
	case "red":
		col = color.RedString
	case "blue":
		col = color.BlueString
	case "yellow":
		col = color.YellowString
	}
	fmt.Fprintln(Out, col(s))
}

// UserInput takes a user input integer
func UserInput(i *int) {
	fmt.Fscan(In, i)
}

// UserInputln takes a user input line
func UserInputln() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n >>> ")
	text, _ := reader.ReadString('\n')
	return text
}
