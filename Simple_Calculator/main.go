package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to a simple CLI Calculator")

	var option int

	result := 0.0

	var Num1, Num2 float64

	for true {
		fmt.Println("\n 1. Addition\n 2. Subtraction\n 3. Multiplication\n 4. Division\n 5. Exit")

		fmt.Print("Select an option: ")

		if _, err := fmt.Scanln(&option); err != nil {
			fmt.Println("Invalid option!")
			continue
		}

		if option <= 0 || option > 5 {
			fmt.Println("Invalid Option!")
			continue
		}

		if option == 5 {
			break
		}

		fmt.Print("Enter the first Num: ")
		if _, err := fmt.Scanln(&Num1); err != nil {
			fmt.Println("Invalid number!")
			continue
		}

		fmt.Print("Enter the Second Num: ")
		if _, err := fmt.Scanln(&Num2); err != nil {
			fmt.Println("Invalid number!")
			continue
		}

		if option == 4 && Num2 == 0 {
			fmt.Println("Cannot divide by zero!")
			continue
		}

		switch option {
		case 1:
			result = Num1 + Num2
			fmt.Printf("%g + %g = %g\n", Num1, Num2, result)
		case 2:
			result = Num1 - Num2
			fmt.Printf("%g - %g = %g\n", Num1, Num2, result)
		case 3:
			result = Num1 * Num2
			fmt.Printf("%g * %g = %g\n", Num1, Num2, result)
		case 4:
			result = Num1 / Num2
			fmt.Printf("%g / %g = %g\n", Num1, Num2, result)

		}
	}

	fmt.Println("Hey Honney, Take care ...!")

}
