package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	fmt.Println("Please enter the animal name and info requested... or enter `q` to terminate.")
	var name, request string
	for true {
		fmt.Printf("> ")
		fmt.Scan(&name)
		if name == "q" || name == "Q" {
			fmt.Println("Thank You :)")
			break
		}
		fmt.Scan(&request)

		var animal Animal
		if name == "cow" {
			animal = Animal{
				"grass",
				"walk",
				"moo",
			}
		} else if name == "bird" {
			animal = Animal{
				"worms",
				"fly",
				"peep",
			}
		} else if name == "snake" {
			animal = Animal{
				"mice",
				"slither",
				"hsss",
			}
		} else {
			fmt.Println("Animal type is not recognized!!")
			continue
		}

		if request == "eat" {
			animal.Eat()
		} else if request == "move" {
			animal.Move()
		} else if request == "speak" {
			animal.Speak()
		} else {
			fmt.Println("Requested info is not recognized.")
		}
	}
}
