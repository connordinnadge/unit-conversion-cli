package main

import (
	"fmt"
	"unit-converter/converter"
	"unit-converter/ui"
)

func main() {

	choice := ui.DisplayMainMenu()

	switch choice {
	case "1":
		converter.LengthConversion()
	case "2":
		converter.WeightConversion()
	case "3":
		converter.VolumeConversion()
	case "4":
		converter.TemperatureConversion()
	default:
		fmt.Println("Invalid choice")
	}
}
