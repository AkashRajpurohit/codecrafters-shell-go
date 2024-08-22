package main

import (
	"bufio"
	"os/exec"
	"strconv"
	"strings"

	// Uncomment this block to pass the first stage
	"fmt"
	"os"
)

var shellBuiltins = []string{"exit", "echo", "type", "pwd"}

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

	paths := os.Getenv("PATH")
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

	case "pwd":
		wd, err := os.Getwd()
		if err != nil {
			fmt.Fprint(os.Stdout, "Error getting working directory\n")
			return
		}

		fmt.Fprintf(os.Stdout, "%s\n", wd)

	case "type":
		if len(arguments) == 0 {
			fmt.Fprint(os.Stdout, "type: missing argument\n")
			return
		}

		command := arguments[0]

		for _, builtin := range shellBuiltins {
			if command == builtin {
				fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", command)
				return
			}
		}

		for _, path := range strings.Split(paths, ":") {
			if _, err := os.Stat(path + "/" + command); err == nil {
				fmt.Fprintf(os.Stdout, "%s is %s/%s\n", command, path, command)
				return
			}
		}

		fmt.Fprintf(os.Stdout, "%s: not found\n", command)
	default:
		for _, path := range strings.Split(paths, ":") {
			if _, err := os.Stat(path + "/" + command); err == nil {
				output, err := exec.Command(path+"/"+command, arguments...).CombinedOutput()

				if err != nil {
					fmt.Fprintf(os.Stdout, "%s\n", err)
					return
				}

				fmt.Fprintf(os.Stdout, "%s", output)
				return
			}
		}

		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}

func main() {
	for {
		readCommand()
	}
}
