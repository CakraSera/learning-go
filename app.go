package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		var menuSelected int
		fmt.Println("Calculator Suhu")
		fmt.Println("Select Menu you want:")
		fmt.Println("1. Celvin to Fahrenheit")
		fmt.Println("4. Exit")
		fmt.Scan(&menuSelected)

		if menuSelected != 4 {
			fmt.Println("menu")
			fmt.Println("You selected menu", menuSelected)
			var input float64
			fmt.Print("Input to convert: ")
			fmt.Scan(&input)
			switch menuSelected {
			case 1:
				var output float64
				output = CelvinToFahrenheit(input)
			}
		} else {
			fmt.Println("good bye")
			return
		}
	}

	fmt.Println("Thanks for using my apps")
}

func CelvinToFahrenheit(input float64) (output float64) {

	fmt.Println("Celvin", input)
	output = (input * 9 / 5) - 459.67
	return
}

func ProfitCalculator() {
	const inflationRate = 2.5

	var investmentAmount float64

	var expectedPercentage, year = 5.5, 10.0

	fmt.Print("Input investment amount: ")

	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow(1+expectedPercentage/100, year)
	futureValueDeflation := investmentAmount * math.Pow(1-inflationRate/100, year)

	fmt.Println("Result after", year, "years is", futureValue)
	fmt.Println("Result deflation after", year, "years is", futureValueDeflation)
}
