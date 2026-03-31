// package main

// import "fmt"

// func main() {

// 	var option string
// 	fmt.Println(`type help to see the help menu`)
// 	fmt.Scan(&option)

// 	if option == "help" {
// 		fmt.Println("type any of the available command")
// 		fmt.Println("add <a>  <b>")
// 		fmt.Println("sub <a>  <b>")
// 		fmt.Println("div <a>  <b>")
// 		fmt.Println("mul <a>  <b>")
// 		fmt.Println("pow <a>  <b>")

// 	}

// }

package main

import (
	"fmt"
	"strconv"
	"strings"
)


func cal() {
start:

	var option string
	fmt.Println("choose your operation")
	fmt.Println("(1) hexToDec")
	fmt.Println("(2) binToDec")
	fmt.Println("(3) DecToBin")
	fmt.Println("(5) help")
	fmt.Scan(&option)
	if option == "quit" {
		fmt.Println("thanks for your service and good bye")
		return
	}

	if option != "1" && option != "2" && option != "3" && option != "help" && option != "quit" {
		fmt.Println("Choose the correct format value")
	}
	if option == "help" {
		fmt.Println("hex convert from decimal to hexadecimal")
		fmt.Println("bin convert from decimal to binary")
		fmt.Println("DecToBin convert from decimal to binary and hexadecimal")
		goto start
	}

	var input string
	fmt.Println("input your value")
	fmt.Scan(&input)
	if option == "quit" {
		fmt.Println("thanks for your service and good bye")
		return

	}
	

	if option == "1" {
		val, err := strconv.ParseInt(input, 16, 64)
		if err != nil {
			fmt.Println("not a valid hex value")
			goto start
		} else {
			fmt.Println("your hextodec answer =", val)
		}
	}
	if option == "2" {
		val, err := strconv.ParseInt(input, 2, 64)
		if err != nil {
			fmt.Println("not a valid bin value")
			goto start
		} else {
			fmt.Println("your binToDec answer =", val)
		}
	}
	if option == "3" {
		mal, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println("not a valid dec value")
			goto start
		} else {
			fmt.Println("bin answer =")
			fmt.Println(strconv.FormatInt(mal, 2))
			fmt.Println("hex answer =")
			fmt.Println(strings.ToUpper(strconv.FormatInt(mal, 16)))
		}

	}

	goto start

}
func runArithmeticCalculator() {
	
	fmt.Println("\n--- Arithmetic Calculator Menu ---")
	fmt.Println("This is where the add, sub, div, mul, pow logic would go.")
	fmt.Println("You would use fmt.Scan to read user input for operation and numbers.")

}

func main() {
	for { 
		var mainOption string
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("Select an application to run:")
		fmt.Println("(C)onverter: Number base conversions")
		fmt.Println("(A)rithmetic: Basic math operations")
		fmt.Println("(Q)uit: Exit the program")
		fmt.Scan(&mainOption)
		
	
		mainOption = strings.ToLower(mainOption)

		if mainOption == "q" || mainOption == "quit" {
			fmt.Println("Thanks for using the combined program. Goodbye!")
			break 
		} else if mainOption == "c" || mainOption == "converter" {
			cal()
		} else if mainOption == "a" || mainOption == "arithmetic" {
			runArithmeticCalculator() 
		} else {
			fmt.Println("Invalid option. Please choose 'C', 'A', or 'Q'.")
		}
	}
}
