package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Type your command: ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			fmt.Println("Empty input")
			continue
		}

		if input == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		parts := strings.Fields(input)

		if len(parts) != 3 || parts[0] != "convert" {
			fmt.Println("Invalid command format. Use: convert value base")
			continue
		}

		value := parts[1]
		base := strings.ToLower(parts[2])

		switch base {
		case "hex":
			handleHex(value)
		case "bin":
			handleBin(value)
		case "dec":
			handleDec(value)
		default:
			fmt.Println("Unsupported base. Use hex, bin, or dec")
		}

		
		fmt.Println("Are you still interested in more conversion? \nType 'yes' to continue or anything else to quit:")
		var name string
		fmt.Scan(&name)

		if name != "yes" {
			fmt.Println("Goodbye!")
			break
		}
	}
}

func handleHex(value string) {
	num, err := strconv.ParseInt(value, 16, 64)
	if err != nil {
		fmt.Println("Invalid hexadecimal number")
		return
	}
	fmt.Printf("Decimal: %d\n", num)
}

func handleBin(value string) {
	for _, ch := range value {
		if ch != '0' && ch != '1' {
			fmt.Println("Invalid binary number")
			return
		}
	}

	num, err := strconv.ParseInt(value, 2, 64)
	if err != nil {
		fmt.Println("Invalid binary number")
		return
	}
	fmt.Printf("Decimal: %d\n", num)
}

func handleDec(value string) {
	num, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		fmt.Println("Invalid decimal number")
		return
	}

	fmt.Printf("Binary:  %b\n", num)
	fmt.Printf("Hex:     %X\n", num)
}