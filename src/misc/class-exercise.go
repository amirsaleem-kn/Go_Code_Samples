package main

import (
	"fmt"
)

// Animal struct
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat method
func (a Animal) Eat() {
	fmt.Println(a.food)
}

// Move method
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

// Speak method
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func handleExec(name, info string) {
	var animal Animal

	if name == "cow" {
		animal = Animal{"grass", "walk", "moo"}
	} else if name == "bird" {
		animal = Animal{"worms", "fly", "peep"}
	} else if name == "snake" {
		animal = Animal{"mice", "slither", "hsss"}
	} else {
		fmt.Println("Invalid Animal", name)
		return
	}

	if info == "speak" {
		animal.Speak()
	} else if info == "eat" {
		animal.Eat()
	} else if info == "move" {
		animal.Move()
	} else {
		fmt.Println("Invalid Info", info)
		return
	}
}

func handleInput() {
	var animalName string
	var animalInfo string

	fmt.Println("> Enter the name of animal and information")
	fmt.Scan(&animalName, &animalInfo)

	handleExec(animalName, animalInfo)

	handleInput()
}

func main() {

	handleInput()

}
