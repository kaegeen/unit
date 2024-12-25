package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Unit Converter")
	fmt.Println("================")
	fmt.Println("Choose a conversion type:")
	fmt.Println("1. Kilometers to Miles")
	fmt.Println("2. Miles to Kilometers")
	fmt.Println("3. Celsius to Fahrenheit")
	fmt.Println("4. Fahrenheit to Celsius")
	fmt.Println("5. Kilograms to Pounds")
	fmt.Println("6. Pounds to Kilograms")
	fmt.Println("Type 'exit' to quit.\n")

	for {
		fmt.Print("Enter choice: ")
		var input string
		fmt.Scanln(&input)

		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > 6 {
			fmt.Println("Invalid choice. Please select a valid option.")
			continue
		}

		fmt.Print("Enter value to convert: ")
		var value float64
		_, err = fmt.Scanln(&value)
		if err != nil {
			fmt.Println("Invalid input. Please enter a numeric value.")
			continue
		}

		switch choice {
		case 1:
			fmt.Printf("%.2f kilometers is %.2f miles\n", value, value*0.621371)
		case 2:
			fmt.Printf("%.2f miles is %.2f kilometers\n", value, value/0.621371)
		case 3:
			fmt.Printf("%.2f Celsius is %.2f Fahrenheit\n", value, (value*9/5)+32)
		case 4:
			fmt.Printf("%.2f Fahrenheit is %.2f Celsius\n", value, (value-32)*5/9)
		case 5:
			fmt.Printf("%.2f kilograms is %.2f pounds\n", value, value*2.20462)
		case 6:
			fmt.Printf("%.2f pounds is %.2f kilograms\n", value, value/2.20462)
		default:
			fmt.Println("Something went wrong. Please try again.")
		}
	}
}
