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


func toUppertoLower(text string) string {
	words := strings.Fields(text)
	result := []string{}

	for i := 0; i < len(words); i++ {
		w := words[i]

		
		if strings.HasPrefix(w, "(") {
			marker := w
		
			for !strings.HasSuffix(marker, ")") && i+1 < len(words) {
				i++
				marker += " " + words[i]
			}

			marker = strings.TrimSuffix(strings.TrimPrefix(marker, "("), ")")
			parts := strings.Split(marker, ",")
			action := strings.TrimSpace(parts[0])
			count := 1
			if len(parts) == 2 {
				if n, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
					count = n
				}
			}

			for j := len(result) - count; j < len(result); j++ {
				if j >= 0 {
					switch action {
					case "up":
						result[j] = strings.ToUpper(result[j])
					case "low":
						result[j] = strings.ToLower(result[j])
					}
				}
			}
			continue
		}

		result = append(result, w)
	}

	return strings.Join(result, " ")
}
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func capitalize(word string) string {
	words := strings.Fields(word)
	result := []string{}

	for i := 0; i < len(words); i++ {
		w := words[i]

		if strings.HasPrefix(w, "(") {
			marker := w
			for !strings.HasSuffix(marker, ")") && i+1 < len(words) {
				i++
				marker += " " + words[i]
			}

			marker = strings.TrimSuffix(strings.TrimPrefix(marker, "("), ")")
			parts := strings.Split(marker, ",")
			action := strings.TrimSpace(parts[0])
			count := 1
			if len(parts) == 2 {
				if n, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
					count = n
				}
			}

			for j := len(result) - count; j < len(result); j++ {
				if j >= 0 {
					switch action {
					case "cap":
						result[j] = capitalizeWord(result[j])
					}
				}
			}
			continue
		}

		result = append(result, w)
	}

	return strings.Join(result, " ")
}

func capitalizeWord(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

