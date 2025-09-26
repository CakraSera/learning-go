package main

import (
	"fmt"
)

func main() {
	for {
		var menuSelected int
		fmt.Println("Calculator Suhu")
		fmt.Println("Select Menu you want:")
		fmt.Println("1. Celvin to Fahrenheit")
		fmt.Println("4. Exit")
		fmt.Scan(&menuSelected)

		if menuSelected == 1 {
			fmt.Println("menu")
		} else {
			fmt.Println("good bye")
			return
		}
	}

	fmt.Println("Thanks for using my apps")
}
