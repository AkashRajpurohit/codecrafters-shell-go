package main

import (
	"bufio"
	"strconv"
	"strings"

	// Uncomment this block to pass the first stage
	"fmt"
	"os"
)

var validCommands = []string{"exit", "echo"}

func readCommand() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	commandString, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
		return
	}

	// Remove newline character
	commandString = strings.TrimSuffix(commandString, "\n")

	// Split the command string into command and arguments
	commandParts := strings.Fields(commandString)
	if len(commandParts) == 0 {
		return
	}

	command := commandParts[0]
	arguments := commandParts[1:]

	for _, validCommand := range validCommands {
		if command == validCommand {
			if command == "exit" {
				exitCode, err := strconv.Atoi(arguments[0])
				if err != nil {
					fmt.Println("Invalid exit code")
					return
				}

				os.Exit(exitCode)
			}

			if command == "echo" {
				fmt.Println(strings.Join(arguments, " "))
			}

			return
		}
	}

	fmt.Printf("%s: command not found\n", command)
}

func main() {
	for {
		readCommand()
	}
}
