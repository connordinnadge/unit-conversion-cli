package converter

import (
	"fmt"
	"math"
	"strconv"
	"unit-converter/ui"
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

func PerformConversion(prompt string, conversionFunc func(float64) float64, inputUnit string, outputUnit string) {
	input := ui.GetUserInput(prompt)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	result := conversionFunc(value)
	ui.DisplayResult(value, inputUnit, result, outputUnit)
}

func WeightConversion() {
	fmt.Println("Please select an option:")
	fmt.Println("1. Pounds to Kilograms")
	fmt.Println("2. Kilograms to Pounds")

	choice := ui.GetUserInput("Enter your choice: ")

	switch choice {
	case "1":
		PerformConversion("Enter the weight in pounds: ", ConvertToKilograms, "pounds", "kilograms")
	case "2":
		PerformConversion("Enter the weight in kilograms: ", ConvertToPounds, "kilograms", "pounds")
	default:
		fmt.Println("Invalid choice")
	}
}

func VolumeConversion() {
	fmt.Println("Please select an option:")
	fmt.Println("1. Gallons to Liters")
	fmt.Println("2. Liters to Gallons")

	choice := ui.GetUserInput("Enter your choice: ")

	switch choice {
	case "1":
		PerformConversion("Enter the volume in gallons: ", ConvertToLiters, "gallons", "liters")
	case "2":
		PerformConversion("Enter the volume in liters: ", ConvertToGallons, "liters", "gallons")
	default:
		fmt.Println("Invalid choice")
	}
}

func LengthConversion() {
	fmt.Println("Please select an option:")
	fmt.Println("1. Kilometers to Miles")
	fmt.Println("2. Miles to Kilometers")
	fmt.Println("3. Centimeters to Inches")
	fmt.Println("4. Inches to Centimeters")

	choice := ui.GetUserInput("Enter your choice: ")

	switch choice {
	case "1":
		PerformConversion("Enter the distance in kilometers: ", ConvertToMiles, "kilometers", "miles")
	case "2":
		PerformConversion("Enter the distance in miles: ", ConvertToKilometers, "miles", "kilometers")
	case "3":
		PerformConversion("Enter the length in centimeters: ", ConvertToInches, "centimeters", "inches")
	case "4":
		PerformConversion("Enter the length in inches: ", ConvertToCentimeters, "inches", "centimeters")
	default:
		fmt.Println("Invalid choice")
	}
}

func TemperatureConversion() {
	fmt.Println("Please select an option:")
	fmt.Println("1. Fahrenheit to Celsius")
	fmt.Println("2. Celsius to Fahrenheit")

	choice := ui.GetUserInput("Enter your choice: ")

	switch choice {
	case "1":
		PerformConversion("Enter the temperature in Fahrenheit: ", ConvertToCelsius, "Fahrenheit", "Celsius")
	case "2":
		PerformConversion("Enter the temperature in Celsius: ", ConvertToFahrenheit, "Celsius", "Fahrenheit")
	default:
		fmt.Println("Invalid choice")
	}
}
