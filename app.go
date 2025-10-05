package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"

	ConvertationTemprature "example.com/profile-calculator/convertation-temprature"
	"github.com/Pallinder/go-randomdata"
)

const balanceFile = "balance.txt"

func main() {
	for {
		var menuSelected int

		fmt.Printf("this data : %d", randomdata.Number(10))
		fmt.Println()
		balance, err := getBalanceFromFile()

		if err != nil {
			fmt.Println("error")
			fmt.Println(err)
			fmt.Println("-----------")
			panic("Can't continue. Sorry")
		}

		fmt.Println("Get Balance", balance)

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
				output = ConvertationTemprature.CelvinToFahrenheit(input)
				writeBalanceToFile(output)
				fmt.Println("Result: ", output)
			}
		} else {
			fmt.Println("good bye")
			return
		}
	}

	fmt.Println("Thanks for using my apps")
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

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFile)

	if err != nil {
		return 1000, errors.New("Failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value")
	}
	return balance, nil
}
