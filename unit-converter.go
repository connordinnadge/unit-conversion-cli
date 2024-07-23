package main

import (
	"fmt"
	"math"
	"strconv"
)

func ConvertToCelsius(fahrenheit float64) float64 {
	return roundToTwoDecimals((fahrenheit - 32) * 5 / 9)
}

func ConvertToFahrenheit(celsius float64) float64 {
	return roundToTwoDecimals(celsius*9/5 + 32)
}

func ConvertToMiles(kilometers float64) float64 {
	return roundToTwoDecimals(kilometers * 0.621371)
}

func ConvertToKilometers(miles float64) float64 {
	return roundToTwoDecimals(miles * 1.60934)
}

func ConvertToKilograms(pounds float64) float64 {
	return roundToTwoDecimals(pounds * 0.453592)
}

func ConvertToPounds(kilograms float64) float64 {
	return roundToTwoDecimals(kilograms * 2.20462)
}

func ConvertToLiters(gallons float64) float64 {
	return roundToTwoDecimals(gallons * 3.78541)
}

func ConvertToGallons(liters float64) float64 {
	return roundToTwoDecimals(liters * 0.264172)
}

func ConvertToInches(centimeters float64) float64 {
	return roundToTwoDecimals(centimeters * 0.393701)
}

func ConvertToCentimeters(inches float64) float64 {
	return roundToTwoDecimals(inches * 2.54)
}

func roundToTwoDecimals(num float64) float64 {
	return math.Round(num*100) / 100
}

func getUserInput(message string) string {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input
}

func main() {

	fmt.Println("Welcome to the Unit Converter!")
	fmt.Println("Please select an option:")
	fmt.Println("1. Length")
	fmt.Println("2. Weight")
	fmt.Println("3. Volume")
	fmt.Println("4. Temperature")

	choice := getUserInput("Enter your choice: ")

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
	input := getUserInput(prompt)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	result := conversionFunc(value)
	fmt.Printf("%v %s is %v %s\n", value, inputUnit, result, outputUnit)
}

func weightConversion() {
	fmt.Println("Please select an option:")
	fmt.Println("1. Pounds to Kilograms")
	fmt.Println("2. Kilograms to Pounds")

	choice := getUserInput("Enter your choice: ")

	switch choice {
	case "1":
		performConversion("Enter the weight in pounds: ", ConvertToKilograms, "pounds", "kilograms")
	case "2":
		performConversion("Enter the weight in kilograms: ", ConvertToPounds, "kilograms", "pounds")
	default:
		fmt.Println("Invalid choice")
	}
}

func volumeConversion() {
	fmt.Println("Please select an option:")
	fmt.Println("1. Gallons to Liters")
	fmt.Println("2. Liters to Gallons")

	choice := getUserInput("Enter your choice: ")

	switch choice {
	case "1":
		performConversion("Enter the volume in gallons: ", ConvertToLiters, "gallons", "liters")
	case "2":
		performConversion("Enter the volume in liters: ", ConvertToGallons, "liters", "gallons")
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

	choice := getUserInput("Enter your choice: ")

	switch choice {
	case "1":
		performConversion("Enter the distance in kilometers: ", ConvertToMiles, "kilometers", "miles")
	case "2":
		performConversion("Enter the distance in miles: ", ConvertToKilometers, "miles", "kilometers")
	case "3":
		performConversion("Enter the length in centimeters: ", ConvertToInches, "centimeters", "inches")
	case "4":
		performConversion("Enter the length in inches: ", ConvertToCentimeters, "inches", "centimeters")
	default:
		fmt.Println("Invalid choice")
	}
}

func temperatureConversion() {
	fmt.Println("Please select an option:")
	fmt.Println("1. Fahrenheit to Celsius")
	fmt.Println("2. Celsius to Fahrenheit")

	choice := getUserInput("Enter your choice: ")

	switch choice {
	case "1":
		performConversion("Enter the temperature in Fahrenheit: ", ConvertToCelsius, "Fahrenheit", "Celsius")
	case "2":
		performConversion("Enter the temperature in Celsius: ", ConvertToFahrenheit, "Celsius", "Fahrenheit")
	default:
		fmt.Println("Invalid choice")
	}
}
