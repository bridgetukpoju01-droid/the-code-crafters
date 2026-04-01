package main

import (
	"regexp"
	"strconv"
	"strings"
)

// Handled by Bridget

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

// Handled by Elizabeth

func Quote(s string) string {
	words := strings.Split(s, "'")
	result := ""
	for i := 0; i < len(words); i++ {
		if i%2 == 1 {
			result += "'" + strings.TrimSpace(words[i]) + "'"
		} else {
			result += words[i]
		}
	}
	return result
}

// Handled by Amina

func BinToDecimal(text string) string {
	words := strings.Fields(text)
	var result []string
	for i := 0; i < len(words); i++ {

		if words[i] == "(bin)" && i > 0 {
			value := words[i-1]
			n, err := strconv.ParseInt(value, 2, 64)
			if err == nil {
				result[len(result)-1] = strconv.FormatInt(n, 10)
			}
			continue
		}
		result = append(result, words[i])
	}
	return strings.Join(result, " ")
}

// Handled by Gift

func fixPunctuation(text string) string {
	reEllipsis := regexp.MustCompile(`\.\s*\.\s*\.`)
	text = reEllipsis.ReplaceAllString(text, "...")

	reBefore := regexp.MustCompile(`\s+([.,!?;:])`)
	text = reBefore.ReplaceAllString(text, "$1")

	reAfter := regexp.MustCompile(`([.,!?;:])([^\s.,!?;:])`)
	text = reAfter.ReplaceAllString(text, "$1 $2")

	reQuotes := regexp.MustCompile(`'\s*(.*?)\s*'`)
	text = reQuotes.ReplaceAllString(text, "'$1'")

	reSpaces := regexp.MustCompile(`\s+`)
	text = reSpaces.ReplaceAllString(text, " ")

	return strings.TrimSpace(text)
}

// Handled by Sunday

func capitalizeWord(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(word[:1]) + word[1:]
}
func getCapCount(token string) int {

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

// Handled by Dominion

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

// Handled by Enock

func articleA(text string) string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return text
	}

	vowels := "aeiouhAEIOUH"
	consonants := "bcdfgjklmnpqrstvwxyzBCDFGJKLMNPQRSTVWXYZ"

	for i := 0; i < len(words)-1; i++ {
		currentWord := strings.ToLower(words[i])
		if len(words[i+1]) == 0 {
			continue
		}

		nextWordFirstChar := string(words[i+1][0])

		if currentWord == "a" && strings.ContainsAny(nextWordFirstChar, vowels) {
			if words[i] == "A" {
				words[i] = "An"
			} else {
				words[i] = "an"
			}
		}

		if currentWord == "an" && strings.ContainsAny(nextWordFirstChar, consonants) {
			if words[i] == "An" {
				words[i] = "A"
			} else {
				words[i] = "a"
			}
		}
	}

	return strings.Join(words, " ")
}
