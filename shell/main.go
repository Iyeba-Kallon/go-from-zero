package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// 1. Show the prompt
		cwd, _ := os.Getwd()
		fmt.Printf("go-shell:%s> ", cwd)

		// 2. Read input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// 3. Clean and parse input
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		args := strings.Fields(input)
		command := args[0]

		// 4. Handle built-in commands
		switch command {
		case "exit":
			fmt.Println("Goodbye!")
			return
		case "cd":
			if len(args) < 2 {
				fmt.Println("Usage: cd <directory>")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Printf("cd error: %v\n", err)
			}
			continue
		}

		// 5. Execute external commands
		cmd := exec.Command(command, args[1:]...)

		// Redirect standard output/error to our terminal
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		err = cmd.Run()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
