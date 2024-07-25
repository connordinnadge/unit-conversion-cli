package ui

import "fmt"

func GetUserInput(message string) string {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input
}

func DisplayMainMenu() string {
	fmt.Println("Welcome to the Unit Converter!")
	fmt.Println("Please select an option:")
	fmt.Println("1. Length")
	fmt.Println("2. Weight")
	fmt.Println("3. Volume")
	fmt.Println("4. Temperature")

	choice := GetUserInput("Enter your choice: ")
	return choice
}

func DisplayResult(fromValue float64, fromUnit string, toValue float64, toUnit string) {
	fmt.Printf("%v %s is %v %s\n", fromValue, fromUnit, toValue, toUnit)
}
