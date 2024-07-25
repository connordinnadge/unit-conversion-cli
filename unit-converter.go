package main

import (
	"fmt"
	"strconv"
	"unit-converter/converter"
	"unit-converter/ui"
)

func main() {

	choice := ui.DisplayMainMenu()

	switch choice {
	case "1":
		lengthConversion()
	case "2":
		weightConversion()
	case "3":
		volumeConversion()
	case "4":
		temperatureConversion()
	default:
		fmt.Println("Invalid choice")
	}
}

func performConversion(prompt string, conversionFunc func(float64) float64, inputUnit string, outputUnit string) {
	input := ui.GetUserInput(prompt)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	result := conversionFunc(value)
	ui.DisplayResult(value, inputUnit, result, outputUnit)
}

func weightConversion() {
	fmt.Println("Please select an option:")
	fmt.Println("1. Pounds to Kilograms")
	fmt.Println("2. Kilograms to Pounds")

	choice := ui.GetUserInput("Enter your choice: ")

	switch choice {
	case "1":
		performConversion("Enter the weight in pounds: ", converter.ConvertToKilograms, "pounds", "kilograms")
	case "2":
		performConversion("Enter the weight in kilograms: ", converter.ConvertToPounds, "kilograms", "pounds")
	default:
		fmt.Println("Invalid choice")
	}
}

func volumeConversion() {
	fmt.Println("Please select an option:")
	fmt.Println("1. Gallons to Liters")
	fmt.Println("2. Liters to Gallons")

	choice := ui.GetUserInput("Enter your choice: ")

	switch choice {
	case "1":
		performConversion("Enter the volume in gallons: ", converter.ConvertToLiters, "gallons", "liters")
	case "2":
		performConversion("Enter the volume in liters: ", converter.ConvertToGallons, "liters", "gallons")
	default:
		fmt.Println("Invalid choice")
	}
}

func lengthConversion() {
	fmt.Println("Please select an option:")
	fmt.Println("1. Kilometers to Miles")
	fmt.Println("2. Miles to Kilometers")
	fmt.Println("3. Centimeters to Inches")
	fmt.Println("4. Inches to Centimeters")

	choice := ui.GetUserInput("Enter your choice: ")

	switch choice {
	case "1":
		performConversion("Enter the distance in kilometers: ", converter.ConvertToMiles, "kilometers", "miles")
	case "2":
		performConversion("Enter the distance in miles: ", converter.ConvertToKilometers, "miles", "kilometers")
	case "3":
		performConversion("Enter the length in centimeters: ", converter.ConvertToInches, "centimeters", "inches")
	case "4":
		performConversion("Enter the length in inches: ", converter.ConvertToCentimeters, "inches", "centimeters")
	default:
		fmt.Println("Invalid choice")
	}
}

func temperatureConversion() {
	fmt.Println("Please select an option:")
	fmt.Println("1. Fahrenheit to Celsius")
	fmt.Println("2. Celsius to Fahrenheit")

	choice := ui.GetUserInput("Enter your choice: ")

	switch choice {
	case "1":
		performConversion("Enter the temperature in Fahrenheit: ", converter.ConvertToCelsius, "Fahrenheit", "Celsius")
	case "2":
		performConversion("Enter the temperature in Celsius: ", converter.ConvertToFahrenheit, "Celsius", "Fahrenheit")
	default:
		fmt.Println("Invalid choice")
	}
}
