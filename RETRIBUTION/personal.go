package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Allhex(value string) string {

	words := strings.Fields(value)

	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			hexValue := words[i-1]
			decimalValue, err := strconv.ParseInt(hexValue, 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(decimalValue, 10)

			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(Allhex("i have 1E (hex) file"))
	fmt.Println(Allhex("ineed FF (hex) mango"))
}
