package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO: Add error handling for user input, i.e. only accept "yes" and "no"

func UserPrompt() (response string) {
	// This is used to ask the user if they want to populate the database or
	// not.
	scanner := bufio.NewReader(os.Stdin)
	fmt.Println(">> Do you want to populate the database? (yes/no): ")
	name, _ := scanner.ReadString('\n')

	return strings.TrimSpace(name)
}
