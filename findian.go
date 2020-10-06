package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a string:")
	var inp string
	if scanner.Scan() {
		inp = scanner.Text()
	}

	foundI := inp[0] == 'i' || inp[0] == 'I'
	foundA := s.Contains(inp, "a") || s.Contains(inp, "A")
	foundN := inp[len(inp)-1] == 'n' || inp[len(inp)-1] == 'N'

	if foundI && foundA && foundN {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
