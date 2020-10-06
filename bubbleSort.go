package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swap(sli []int, idx int) {
	var temp int
	temp = sli[idx]
	sli[idx] = sli[idx+1]
	sli[idx+1] = temp
}

/*
BubbleSort function to perform the bubble sort algorithm
*/
func BubbleSort(sli []int) {
	n := len(sli)
	for i := 0; i < n-1; i++ {
		for j := 0; j+i+1 < n; j++ {
			if sli[j] > sli[j+1] {
				swap(sli, j)
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sli := make([]int, 0)
	fmt.Println("Enter the sequence of intergers seprated by space:")
	if scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		for _, val := range input {
			if len(val) == 0 {
				continue
			}
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			sli = append(sli, num)
		}
	}
	//fmt.Println(sli)

	BubbleSort(sli)

	fmt.Println("Sorted integer sequence...")
	fmt.Println(sli)

}

/*
Test cases:
5 1 4 2 8
64 34 25 12 22 11 -90
12 34 65 -0 09
*/
