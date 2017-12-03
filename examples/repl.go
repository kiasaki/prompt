package main

import (
	"fmt"
	"strings"

	"github.com/kiasaki/prompt"
)

var P = prompt.NewPrompt()

func read() (string, bool) {
	line, err := P.Prompt("csv> ")
	return line, err == nil
}

func eval(line string) []string {
	P.AppendHistory(line)
	return strings.Split(line, ",")
}

func print(cells []string) {
	for i, cell := range cells {
		fmt.Println(i+1, ": "+cell)
	}
}

func main() {
	fmt.Println("'CSV' parser. Press Ctrl-D to exit.")
	for {
		if line, ok := read(); ok {
			print(eval(line))
		} else {
			fmt.Println()
			break
		}
	}
}
