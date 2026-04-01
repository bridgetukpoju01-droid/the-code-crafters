package main

import (
	
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

func Quote(s string) string {
	words := strings.Split(s, "'") 
	result := "  "
	for i := 0; i < len(words); i++ {
		if i%2 == 1 {
		   result += "'" + strings.TrimSpace(words[i]) + "'" 
		} else {
			result += words[i]
		}
	}
	return result
}


func ToUpper(input string) string {
    t := "(up)"

    words := strings.Fields(input)

    for i := 0; i < len(words); i++ {
        w := words[i]

        switch {
        case w == "(up)":

            if w == t && i > 0 {
                words[i-1] = strings.ToUpper(string(words[i-1]))
                words = append(words[:i], words[i+1:]...)

            }

        case strings.HasPrefix(words[i], "(up,"):

            num := strings.TrimRight(words[i+1], ")")
            n, _ := strconv.Atoi(num)

            for j := i - n; j < i; j++ {
                words[j] = strings.ToUpper(string(words[j]))
                words = append(words[:i], words[i+1:]...)
            }

        }

    }
    return strings.Join(words, " ")

}

func ToLower(input string) string {
    t := "(low)"

    words := strings.Fields(input)

    for i := 0; i < len(words); i++ {
        w := words[i]

        switch {
        case w == "(low)":

            if w == t && i > 0 {
                words[i-1] = strings.ToLower(string(words[i-1]))
                words = append(words[:i], words[i+1:]...)

            }

        case strings.HasPrefix(words[i], "(low,"):

            num := strings.TrimRight(words[i+1], ")")
            n, _ := strconv.Atoi(num)

            for j := i - n; j < i; j++ {
                words[j] = strings.ToLower(string(words[j]))
                words = append(words[:i], words[i+1:]...)
            }

        }

    }
    return strings.Join(words, " ")

}