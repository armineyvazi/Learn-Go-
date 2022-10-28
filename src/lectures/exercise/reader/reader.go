//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:

//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	CmdHello = "hello"
	CmdBye   = "bye"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numLines := 0
	numCommands := 0

	for scanner.Scan() {
		//* When the user enters either "hello" or "bye", the program
		if strings.ToUpper(scanner.Text()) == "Q" {
			break
		} else {
			text := strings.TrimSpace(scanner.Text())
			//  should respond with a message after pressing the enter key.
			//* Whenever the user types a "Q" or "q", the program should exit.
			switch text {
			case CmdHello:
				numCommands += 1
				fmt.Println("command respose:Hi")
			case CmdBye:
				numCommands += 1
				fmt.Println("command response:Bye")
			}
			if text != "" {
				numLines += 1
			}

		}
	}

	fmt.Printf("You entered %v lines\n", numLines)
	fmt.Printf("You entered %v commands\n", numCommands)

}
