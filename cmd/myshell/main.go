package main

import (
	"bufio"
	// Uncomment this block to pass the first stage
	"fmt"
	"os"
)

var validCommands = []string{}

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
		return
	}

	// Remove newline character
	command = command[:len(command)-1]

	for _, validCommand := range validCommands {
		if command == validCommand {
			fmt.Println("Command found")
			return
		}
	}

	fmt.Printf("%s: command not found\n", command)
}
