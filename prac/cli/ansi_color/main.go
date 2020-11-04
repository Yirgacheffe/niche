package main

import "fmt"

// Color of text, assign to int
type Color int

// ColorNone as default
const (
	ColorNone = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	Black Color = -1
)

// ColorText hold a text and it's color
type ColorText struct {
	TextColor Color
	Text      string
}

func (r *ColorText) String() string {

	if r.TextColor == ColorNone {
		return r.Text
	}

	value := 30

	if r.TextColor != Black {
		value += int(r.TextColor)
	}

	return fmt.Sprintf("\033[%dm%s\033[0m", value, r.Text)

}

func main() {

	r := ColorText{Red, "I'm red!"}
	fmt.Println(r.String())

	r.TextColor = Green
	r.Text = "I'm green!"
	fmt.Println(r.String())

	r.TextColor = ColorNone
	r.Text = "Back to normal color..."

	fmt.Println(r.String())

}
