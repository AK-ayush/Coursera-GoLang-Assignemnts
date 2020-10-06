package main

import "fmt"

// Animal is an interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

type cow struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (c cow) Eat() {
	fmt.Println(c.food)
}

func (c cow) Move() {
	fmt.Println(c.locomotion)
}

func (c cow) Speak() {
	fmt.Println(c.noise)
}

type bird struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (b bird) Eat() {
	fmt.Println(b.food)
}

func (b bird) Move() {
	fmt.Println(b.locomotion)
}

func (b bird) Speak() {
	fmt.Println(b.noise)
}

type snake struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (s snake) Eat() {
	fmt.Println(s.food)
}

func (s snake) Move() {
	fmt.Println(s.locomotion)
}

func (s snake) Speak() {
	fmt.Println(s.noise)
}

func main() {
	fmt.Println("Please enter the animal name and info requested... or enter `q` to terminate.")
	var command, name, typ, request string

	cowMap := make(map[string]cow)
	birdMap := make(map[string]bird)
	sankeMap := make(map[string]snake)

	for true {
		fmt.Printf(">")
		fmt.Scan(&command)
		fmt.Scan(&name)

		if command == "newanimal" {
			fmt.Scan(&typ)
			if typ == "cow" {
				cowMap[name] = cow{
					name,
					"grass",
					"walk",
					"moo",
				}
			} else if typ == "bird" {
				birdMap[name] = bird{
					name,
					"worms",
					"fly",
					"peep",
				}
			} else if typ == "snake" {
				sankeMap[name] = snake{
					name,
					"mice",
					"slither",
					"hsss",
				}
			} else {
				fmt.Println("Please enter a valid animal type!")
				continue
			}
			fmt.Println("Created it!")
		} else if command == "query" {
			fmt.Scan(&request)
			var a Animal

			_, isCow := cowMap[name]
			_, isBird := birdMap[name]
			_, isSnake := sankeMap[name]

			if isCow {
				a = cowMap[name]
			} else if isBird {
				a = birdMap[name]
			} else if isSnake {
				a = sankeMap[name]
			} else {
				fmt.Println("Animal with the requested name is not present!")
				continue
			}

			if request == "eat" {
				a.Eat()
			} else if request == "move" {
				a.Move()
			} else if request == "speak" {
				a.Speak()
			} else {
				fmt.Println("Please enter  a valid request for the animal!")
				continue
			}

		} else {
			fmt.Println("Please enter a valid command!")
			continue
		}
	}
}
