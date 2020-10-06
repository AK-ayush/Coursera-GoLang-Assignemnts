package main

import "fmt"

func main() {
	fmt.Println( "Enter a floting point number: " )
	var num float32
	fmt.Scan(&num)
	intNum := int32(num)
	fmt.Printf( "Truncater integer: %d \n", intNum)
}