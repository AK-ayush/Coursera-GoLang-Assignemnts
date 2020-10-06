package main

import (
	"fmt"
)

/*
GenDisplaceFn is function generate the
*/
func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return (0.5)*a*t*t + v0*t + s0
	}
}

func main() {
	var a, v0, s0, t float64
	fmt.Println("Enter the initial values...")

	fmt.Printf("acceleration : ")
	fmt.Scanln(&a)

	fmt.Printf("initial velocity : ")
	fmt.Scanln(&v0)

	fmt.Printf("intial displacement : ")
	fmt.Scanln(&s0)

	fn := GenDisplaceFn(a, v0, s0)
	for true {
		fmt.Println("Enter the value of time : \nOr enter -1 to quit!")
		fmt.Scanln(&t)
		if t == -1 || t < 0 {
			break
		}
		fmt.Printf("displacement : ")
		fmt.Println(fn(t))
	}
	fmt.Println("Thank you!")
}
