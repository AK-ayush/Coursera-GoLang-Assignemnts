package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var name, address string

	fmt.Println("Enter name : ")
	if scanner.Scan() {
		name = scanner.Text()
	}

	fmt.Println("Enter address : ")
	if scanner.Scan() {
		address = scanner.Text()
	}

	idMap := map[string]string{
		"name":    name,
		"address": address,
	}

	outJSON, err := json.MarshalIndent(idMap, "", "\t")
	if err == nil {
		jsonStr := string(outJSON)
		fmt.Println("--------------\nOutput JSON : ")
		fmt.Println(jsonStr)
	}

}
