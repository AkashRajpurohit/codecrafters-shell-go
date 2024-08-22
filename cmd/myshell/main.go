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

	case "type":
		if len(arguments) == 0 {
			fmt.Fprint(os.Stdout, "type: missing argument\n")
			return
		}

		command := arguments[0]
		paths := os.Getenv("PATH")

		// search for the command in the PATH
		for _, path := range strings.Split(paths, ":") {
			if _, err := os.Stat(path + "/" + command); err == nil {
				fmt.Fprintf(os.Stdout, "%s is %s/%s\n", command, path, command)
				return
			}
		}

		fmt.Fprintf(os.Stdout, "%s: not found\n", command)
	default:
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}

func main() {
	for {
		readCommand()
	}
}
