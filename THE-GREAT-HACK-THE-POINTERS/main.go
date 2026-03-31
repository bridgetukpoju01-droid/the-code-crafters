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
	"bufio"
	"os"
)


var command string

func str() {
	fmt.Print("Enter a text: ")
	line := bufio.NewScanner(os.Stdin)
	if line.Scan() {
		line := line.Text()
		if len(line) > 0 {
			fmt.Println("Input Operation!")
			var operation string
			_, err := fmt.Scan(&operation)
			if err == nil {
				if operation == "upper" {
					fmt.Println(strings.ToUpper(line))
				} else if operation == "lower" {
					fmt.Println(strings.ToLower(line))
				} else if operation == "cap" {
					nval := strings.Fields(line)
					for i := 0; i <= len(nval)-1; i++ {
						low := strings.ToLower(nval[i])
						title := strings.ToUpper(string(low[0])) + string(low[1:])
						fmt.Print(title, " ")
						
					}
				} else if operation == "title" {
					var small = []string{"a", "an", "the", "and", "but",
						"or", "for", "nor", "on", "at", "to", "by", "in",
						"of", "up", "as", "is", "it"}
					nval := strings.Fields(line)
					for i := 0; i <= len(nval)-1; i++ {
						for j := 0; j <= len(small)-1; j++ {
							
						}
					}
				} else {
					fmt.Println("Invalid Operation! ")
				}
			} else {
				fmt.Println("An error occured during scan")
			}
		} else {
			fmt.Println("Does not accept an empty input")
		}
	} else {
		fmt.Println("An error occured while reading input!")
	}
}







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
		fmt.Println("(C)converter: Number base conversions")
		fmt.Println("(A)arithmetic: Basic math operations")
		fmt.Println("(s)string converter")
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
		} else if mainOption == "s" {
			str()
		}else {
			fmt.Println("Invalid option. Please choose 'C', 'A', or 'Q'.")
		}
	}
}




package main

import (
	"fmt"
	"strconv"
)

var num1 string
var num2 string
var operation string
var exit int

func addition() {

	fmt.Print("Choose the first digit: ") // Here
start:
	fmt.Scanln(&num1)
	val1, err := strconv.Atoi(num1)
	if err == nil {
		fmt.Print("Input the second digit: ")
	start2:
		fmt.Scanln(&num2)
		val2, err := strconv.Atoi(num2)
		if err == nil {
			fmt.Printf("%v + %v = %v\n", val1, val2, val1+val2)
		} else {
			fmt.Println("Select a valid digit!")
			goto start2
		}
	} else {
		fmt.Println("Invalid input, Choose a digit!")
		goto start
	}

}

func subtraction() {
	fmt.Print("Choose the first digit: ")
start1:
	fmt.Scan(&num1)
	val1, err := strconv.Atoi(num1)
	if err == nil {
		fmt.Print("Input the second digit: ")
	start2:
		fmt.Scanln(&num2)
		val2, err := strconv.Atoi(num2)
		if err == nil {
			fmt.Printf("%v - %v = %v\n", val1, val2, val1-val2)
		} else {
			fmt.Print("Input a valid digit: ")
			goto start2
		}
	} else {
		fmt.Print("Input a valid digit: ")
		goto start1
	}
}

func division() {
	fmt.Print("Choose a first digit: ")
start:
	fmt.Scan(&num1)
	val1, err := strconv.ParseFloat(num1, 64)
	if err == nil {
		fmt.Print("Choose a second digit: ")
	start1:
		fmt.Scan(&num2)
		val2, err := strconv.ParseFloat(num2, 64)
		if err == nil && val2 != 0 {
			fmt.Printf("%v / %v = %v\n", val1, val2, val1/val2)
		} else {
			fmt.Print("Input a valid digit, Zero is not allowed: ")
			goto start1
		}
	} else {
		fmt.Print("Input a valid digit: ")
		goto start
	}

}

func multiplication() {
	fmt.Print("Choose a first digit: ")
start1:
	fmt.Scan(&num1)
	val1, err := strconv.Atoi(num1)
	if err == nil {
		fmt.Print("Choose a second digit: ")
	start2:
		fmt.Scan(&num2)
		val2, err := strconv.Atoi(num2)
		if err == nil {
			fmt.Printf("%v * %v = %v\n", val1, val2, val1*val2)
		} else {
			fmt.Print("Input a valid digit: ")
			goto start2
		}
	} else {
		fmt.Print("Input a valid digit: ")
		goto start1
	}
}

func main() {
start0:
fmt.Println(" ")
	fmt.Println("Choose an operation:\n1. addition\n2. subtraction\n3. multiplication\n4. division\n5. Help\n6. Exit")
	fmt.Println("Choose \"Help\" for suggestions! ")
	fmt.Print(">>> ")
	for i := 0; i <= 6; i++ {
	start1:
		fmt.Scanln(&operation)
		opps, err := strconv.Atoi(operation)
		if err == nil {
			if opps == 1 {
				addition()
				goto start0
			} else if opps < 1 || opps > 6 {
				fmt.Print("Invalid Operation\n\n")
				goto start0
			} else if opps == 2 {
				subtraction()
				goto start0
			} else if opps == 3 {
				multiplication()
				goto start0
			} else if opps == 4 {
				division()
				goto start0
			} else if opps == 5 {
				fmt.Print("Choose any of the following for an operation:\n1. To perform an addition\n2. To perform a subtraction\n3. To perform a multiplication\n4. To perform a division\n5. Help\n6. Exit\n")
				goto start1
			} else if opps == 6 {
				var res string
				fmt.Print("Do you want to exit the program? (y/n): ")
				start2:
				fmt.Scanln(&res)
				if res == "y" || res == "Y" {
					fmt.Println("Goodbye Codecrafter!")
					break
				} else if res == "n" || res == "N"{
					goto start0
				} else {
					fmt.Print("Input a valid response: ")
					goto start2
				}
			}
		} else {
			fmt.Print("Input a valid operation, or choose \"Help\" for assistance: ")
			goto start1
		}

	}
}





