package converter

import "math"

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
