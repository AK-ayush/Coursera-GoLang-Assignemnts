package main

import (
	"fmt"
	"sort"
	s "strconv"
)

func main() {
	sli := make([]int, 3)
	for true {
		var inp string
		fmt.Println("Enter an integer: ")
		fmt.Scan(&inp)
		if inp == "X" {
			break
		}
		num, err := s.Atoi(inp)
		if err != nil {
			fmt.Println("Not a valid integer!!")
			continue
		}
		sli = append(sli, num)
		sort.Slice(sli, func(i, j int) bool { return sli[i] < sli[j] })
		fmt.Printf("Sorted Slice : ")
		fmt.Println(sli)
	}
	fmt.Println("Thanks, See you again!")
}
