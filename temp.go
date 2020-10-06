package main

import (
	"fmt"
	"strings"
)

func main() {
	findian()
}

func findian() {
	var myString string
	fmt.Scan(&myString)
	myString = strings.ToLower(myString)
	if beginWithI(myString) && containA(myString) && endWithN(myString) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}

func beginWithI(str string) bool {
	return str[0] == 'i'
}

func containA(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] == 'a' {
			return true
		}
	}
	return false
}

func endWithN(str string) bool {
	return str[len(str)-1] == 'n'
}
