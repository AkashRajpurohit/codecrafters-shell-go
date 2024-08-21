package main

import (
	"bufio"
	"strconv"
	"strings"

	// Uncomment this block to pass the first stage
	"fmt"
	"os"
)

func readCommand() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	commandString, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
		return
	}

	// Remove newline character
	commandString = strings.TrimSpace(commandString)

	// Split the command string into command and arguments
	commandParts := strings.Fields(commandString)
	if len(commandParts) == 0 {
		return
	}

	command := commandParts[0]
	arguments := commandParts[1:]

	switch command {
	case "exit":
		exitCode, err := strconv.Atoi(arguments[0])
		if err != nil {
			fmt.Fprint(os.Stdout, "Invalid exit code")
			return
		}

		os.Exit(exitCode)

	case "echo":
		fmt.Fprintf(os.Stdout, "%s\n", strings.Join(arguments, " "))

	default:
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}

func main() {
	for {
		readCommand()
	}
}
