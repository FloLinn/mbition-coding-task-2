package main

// This code is not asked for by the coding task.
// I just had fun learning go here :)
// I just presents some test cases to the user
// you can ignore this

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const embossColor = "\033[1;34m%s\033[0m"

// UI is a simple ui interface (just for fun)
type UI struct {
	reader     *bufio.Reader
	MainChoice Choice
	Count      int
}

// Choice enum
type Choice int

const (
	// DEFAULT intervals
	DEFAULT Choice = iota + 1
	// CUSTOM sized interval set
	CUSTOM
	// INFINITE stream of intervals
	INFINITE
)

// GetUserChoices fill UI struct with user input
func (ui *UI) GetUserChoices() {
	fmt.Printf(embossColor, "hello mbition\n")
	fmt.Println("Choose a set of intervals to be merged")
	fmt.Printf(embossColor, "1   ")
	fmt.Println("for your example set ([25,30] [2,19] [14, 23] [4,8])")
	fmt.Printf(embossColor, "2   ")
	fmt.Println("custom sized")
	fmt.Printf(embossColor, "3   ")
	fmt.Println("infinite stream (throttled)")
	validateMainChoice := func(s string) bool {
		i, err := strconv.Atoi(s)
		return err == nil && i >= 1 && i <= 3
	}
	s := ui.getSingleChoice("your choice: ", validateMainChoice)
	mainChoice, _ := strconv.Atoi(s)
	ui.MainChoice = Choice(mainChoice)
	if ui.MainChoice == CUSTOM {
		fmt.Println("how many intervals?")
		validateCount := func(s string) bool {
			i, err := strconv.Atoi(s)
			return err == nil && i >= 1
		}
		c := ui.getSingleChoice("count: ", validateCount)
		count, _ := strconv.Atoi(c)
		ui.Count = count
	}
}

func (ui UI) getSingleChoice(prompt string, validate func(string) bool) string {
	done := false
	text := ""
	var err error
	for !done {
		fmt.Print(prompt)
		text, err = ui.reader.ReadString('\n')
		if err == nil {
			text = strings.TrimSuffix(text, "\n")
			if validate(text) {
				done = true
			}
		}
		if !done {
			fmt.Println("try again")
		}
	}
	return text
}

// NewUI creates simple UI for merger options
func NewUI() UI {
	return UI{bufio.NewReader(os.Stdin), DEFAULT, 0}
}
