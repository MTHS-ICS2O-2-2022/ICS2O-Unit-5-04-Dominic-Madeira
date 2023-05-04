// Created by: Dominic M.
// Created on: April 2023
//
// This program checks to see if you are eligible for student pricing.
package main

import (
	"fmt"
)

func main() {

	var day string
	var age int

	// input
	fmt.Println("This program checks to see if you are eligible for student pricing.")
	fmt.Println()
	fmt.Print("Day is it? ")
	fmt.Println()
	fmt.Scanln(&day)
	fmt.Println()
	fmt.Print("How old are you? ")
	fmt.Println()
	fmt.Scanln(&age)

	// process
	if day == "tuesday" || day == "thursday" || day == "Tuesday" || day == "Thursday" || age > 12 && age < 21 {
		fmt.Println("You are eligible for student pricing.")
		fmt.Println()
		fmt.Print("\nDone.")
		fmt.Println()
	} else {
		fmt.Println("You are not eligible for student pricing.")
		fmt.Println()
		fmt.Print("\nDone.")
		fmt.Println()

	}
}
