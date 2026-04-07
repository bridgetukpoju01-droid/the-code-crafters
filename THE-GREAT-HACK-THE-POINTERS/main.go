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
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var command string

func str() {
	fmt.Print("Enter a text: ")
	line := bufio.NewScanner(os.Stdin)
	if line.Scan() {
		line := line.Text()
		if len(line) > 0 {
			fmt.Print("Select Operation!\n1. Uppercase\n2. Lowercase\n3. Capitalize\n4. Title Case\n5. Snakecase\n6. Reverse\n>>>")
			var operation string
			_, err := fmt.Scan(&operation)
			if err == nil {
				if operation == "1" {
					fmt.Println(strings.ToUpper(line))
				} else if operation == "2" {
					fmt.Println(strings.ToLower(line))
				} else if operation == "3" {
					nval := strings.Fields(line)
					for i := 0; i <= len(nval)-1; i++ {
						low := strings.ToLower(nval[i])
						title := strings.ToUpper(string(low[0])) + string(low[1:])
						fmt.Print(title, " ")
					}
				} else if operation == "4" {
					var smallWords = []string{"a", "an", "the", "and", "but",
						"or", "for", "nor", "on", "at", "to", "by",
						"in", "of", "up", "as", "is", "it"}
					words := strings.Fields(line)

					for i, word := range words {
						isSmall := slices.Contains(smallWords, strings.ToLower(word))
						if i == 0 || !isSmall {
							words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
						} else {
							words[i] = strings.ToLower(word)
						}
					}
					fmt.Println(strings.Join(words, " "))
				} else if operation == "5" {
					line = strings.ToLower(line)
					result := ""
					for i := 0; i < len(line); i++ {
						letter := line[i] >= 'a' && line[i] <= 'z'
						digit := line[i] >= '0' && line[i] <= '9'
						space := line[i] == ' '
						if letter || digit {
							result += string(line[i])
						} else if space {
							result += "_"
						}
					}
					for strings.Contains(result, "__") {
						result = strings.ReplaceAll(result, "__", "_")
					}
					result = strings.Trim(result, "_")
					fmt.Println(result)
				} else if operation == "6" {
					words := strings.Fields(line)
					for i := 0; i < len(words); i++ {
						word := words[i]
						for j := len(word) - 1; j >= 0; j-- {
							fmt.Print(string(word[j]))
						}
						fmt.Print(" ")
					}
					fmt.Println()
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
		fmt.Println("(S)string converter")
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
		} else if mainOption == "s" || mainOption == "S" {
			str()
		} else {
			fmt.Println("Invalid option. Please choose 'C', 'A', or 'Q'.")
		}
	}
}
