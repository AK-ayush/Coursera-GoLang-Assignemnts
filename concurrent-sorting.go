package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sortChunk(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Sorting chunk : ", arr)
	if len(arr) > 1 {
		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	}

}

func mergeChunk(arr1 []int, arr2 []int, arr []int, wg *sync.WaitGroup) {

	defer wg.Done()
	n := len(arr1)
	m := len(arr2)

	i, j, k := 0, 0, 0

	for i < n && j < m {
		if arr1[i] < arr2[j] {
			arr[k] = arr1[i]
			i++
		} else {
			arr[k] = arr2[j]
			j++
		}
		k++
	}

	for i < n {
		arr[k] = arr1[i]
		k++
		i++
	}

	for j < m {
		arr[k] = arr2[j]
		k++
		j++
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var wg sync.WaitGroup
	arr := make([]int, 0)

	fmt.Println("Enter the sequence of integers separated by space:")
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
			arr = append(arr, num)
		}
	}

	n := len(arr)
	if n == 0 {
		fmt.Printf("Sorted array : ")
		fmt.Println(arr)
		return
	}

	// Creating four parts out of input integer seequence
	part1 := arr[0 : n/4]
	var part2, part3, part4 []int

	if n/4 < n {
		part2 = arr[n/4 : min(n/2, n)]
	}
	if n/2 < n {
		part3 = arr[n/2 : min(n, int((3*n)/4))]
	}
	if (3*n)/4 < n {
		part4 = arr[(3*n)/4 : n]
	}

	// fmt.Println(arr)
	// fmt.Println(part1, part2, part3, part4)

	// Invoking goroutines to sort each part individually
	wg.Add(4)
	go sortChunk(part1, &wg)
	go sortChunk(part2, &wg)
	go sortChunk(part3, &wg)
	go sortChunk(part4, &wg)

	// waiting for the above goroutines to finish
	wg.Wait()

	// fmt.Println(part1, part2, part3, part4)

	wg.Add(2)
	part12 := make([]int, len(part1)+len(part2))
	part34 := make([]int, len(part3)+len(part4))

	// merge the part1 and part2 --> part12
	go mergeChunk(part1, part2, part12, &wg)

	// merge the part3 and part4 --> part34
	go mergeChunk(part3, part4, part34, &wg)

	// waiting for the above goroutines to finish
	wg.Wait()

	wg.Add(1)
	sortedArr := make([]int, len(part12)+len(part34))

	// merge the part12 and part34 --> sortedArr
	go mergeChunk(part12, part34, sortedArr, &wg)
	wg.Wait()

	fmt.Println("Sorted array : ", sortedArr)
}
