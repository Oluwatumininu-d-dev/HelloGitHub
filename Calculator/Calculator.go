package Calculator

import (
	"fmt"
)

// FUNCTION TO FIND SUM
func Sum(num1, num2 float64) float64 {
	return num1 + num2
}

// FUNCTION TO SUBTRACT
func Subtract(num1, num2 float64) float64 {
	return num1 - num2
}

// FUNCTION TO MULTIPLY
func Multiply(num1, num2 float64) float64 {
	return num1 * num2
}

// FUNCTION TO DIVIDE
func Divide(num1, num2 float64) float64 {
	if num2 == 0 {
		return 0
	}
	return num1 / num2
}

// FUNCTION TO DISPLAY OPTION
func DisplayOption() {
	fmt.Println("1. Sum")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("----------------")
}

func main() {
	var Again string

	for Again != "Q" && Again != "q" {
		fmt.Println("")
		fmt.Println("")
		fmt.Println("ENTER A NUMBER TO PERFORM A CALCULATION")
		DisplayOption()

		var Input int
		var Num1 float64
		var Num2 float64

		fmt.Print("WHAT OPERATION DO YOU WANT TO PERFORM: \t")
		_, InputErr := fmt.Scan(&Input)
		if InputErr != nil {
			fmt.Println(InputErr)
		}

		fmt.Print("ENTER NUMBER 1: \t")
		_, Num1Err := fmt.Scan(&Num1)
		if Num1Err != nil {
			fmt.Println(Num1Err)
		}

		fmt.Print("ENTER NUMBER 2: \t")
		_, Num2Err := fmt.Scan(&Num2)
		if Num2Err != nil {
			fmt.Println(Num2Err)
		}

		fmt.Println("----------------")

		switch Input {
		case 1:
			sum := Sum(Num1, Num2)
			fmt.Printf("The sum of %f and %f is %f \n", Num1, Num2, sum)
		case 2:
			subtract := Subtract(Num1, Num2)
			fmt.Printf("The subtraion of %f and %f is %f \n", Num1, Num2, subtract)
		case 3:
			multiply := Multiply(Num1, Num2)
			fmt.Printf("The Product of %f and %f is %f \n", Num1, Num2, multiply)
		case 4:
			divide := Divide(Num1, Num2)
			fmt.Printf("The sum of %f and %f is %f \n", Num1, Num2, divide)
		default:
			fmt.Println("Unknown number specified")
		}

		fmt.Println("Enter any key to perform anoter operatioon and enter 'Q' to quit")
		_, AgainErr := fmt.Scan(&Again)
		if AgainErr != nil {
			fmt.Println(AgainErr)
		}
	}
}
