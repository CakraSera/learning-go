package ConvertationTemprature

import "fmt"

func CelvinToFahrenheit(input float64) (output float64) {
	fmt.Println("Celvin", input)
	output = (input * 9 / 5) - 459.67
	return
}
