package main

import (
	"fmt"
)

func addition(a, b float64) float64 {
	return a + b
}
func division(a, b float64) float64 {
	return a / b
}
func multiplication(a, b float64) float64 {
	return a * b
}
func subtraction(a, b float64) float64 {
	return a - b

}
func main() {
	for {
		var a, b float64
		var operation string
		fmt.Println("choose what you want to do now.")
		fmt.Println("1: addition")
		fmt.Println("2: division")
		fmt.Println("3: multiplication")
		fmt.Println("4: subtraction")
		fmt.Println("5: quit")
		fmt.Println("6:help")
		fmt.Scanln(&operation)
		if operation == "5" {
			fmt.Println("Goodbye")
			break
		}

		switch operation {

		case "1":

			fmt.Println("Enter the first number")
			fmt.Scanln(&a)
			fmt.Println("Enter the second number")
			fmt.Scanln(&b)
			answer := addition(a, b)
			fmt.Println(answer)

		case "2":
			fmt.Println("Enter the first number")
			fmt.Scanln(&a)
			fmt.Println("Enter the second number")
			fmt.Scanln(&b)
			if b != 0 {
				fmt.Println("a,b")
			}
			if b == 0 {
				fmt.Println("cannot divide by 0 :")
			}
		case "3":
			fmt.Println("Enter the first number")
			fmt.Scanln(&a)
			fmt.Println("Enter the second number")
			fmt.Scanln(&b)
			answer := multiplication(a, b)
			fmt.Println(answer)

		case "4":
			fmt.Println("Enter the first number")
			fmt.Scanln(&a)
			fmt.Println("Enter the second number")
			fmt.Scanln(&b)
			answer := subtraction(a, b)
			fmt.Println(answer)

		case "6":
			fmt.Println("1: addition,2:division,3:multiplication,4:subtraction,5:quit")

		}
	}
}
