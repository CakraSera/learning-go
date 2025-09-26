package GoBasic

import (
	"fmt"
	"math"
)

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
