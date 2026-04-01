package main

import (
	"strings"
)

func capitalizeWord(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(word[:1]) + word[1:]
}
func getCapCount(token string) int {
	l
	if token == "(cap)" {
		return 1
	}

	token = strings.TrimPrefix(token, "(cap,")
	token = strings.TrimSuffix(token, ")")

	n := 0
	for _, ch := range token {
		if ch >= '0' && ch <= '9' {
			n = n*10 + int(ch-'0')
		}
	}

	if n <= 0 {
		return 1
	}
	return n
}

func capitalize(s string) string {
	words := strings.Fields(s)
	result := []string{}

	for _, w := range words {
		if strings.HasPrefix(w, "(cap") {
			count := getCapCount(w)

			for i := 0; i < count && i < len(result); i++ {
				idx := len(result) - 1 - i
				result[idx] = capitalizeWord(result[idx])
			}
		} else {
			result = append(result, w)
		}
	}

	return strings.Join(result, " ")
}
