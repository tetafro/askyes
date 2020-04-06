package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const usage = `Asks yes/no question.

Supported answers (case insensitive):
    - yes/no
    - y/n

Usage:
    askyes [QUESTION]`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(0)
	}
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println(usage)
		os.Exit(0)
	}

	question := strings.Join(os.Args[1:], " ")
	fmt.Print(question + " ")

	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Unexpected error: %v", err)
		os.Exit(1)
	}
	answer = strings.ToLower(strings.TrimSpace(answer))

	switch answer {
	case "yes", "y":
		os.Exit(0)
	case "no", "n":
		os.Exit(1)
	default:
		fmt.Println("Unknown type of answer")
		os.Exit(1)
	}

}
