// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [bridget ukpoju]
// Squad:  [the pointers]

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func toUpper(text string) string {
	return strings.ToUpper(text)
}

func toLower(text string) string {
	return strings.ToLower(text)
}

func capitalize(text string) string {
	words := strings.Fields(text)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])
		}
	}
	return strings.Join(words, " ")
}

func titleCase(text string) string {
	small := map[string]bool{
		"a": true, "an": true, "the": true, "and": true, "but": true,
		"or": true, "for": true, "nor": true, "on": true, "at": true,
		"to": true, "by": true, "in": true, "of": true, "up": true,
		"as": true, "is": true, "it": true,
	}
	words := strings.Fields(strings.ToLower(text))
	for i, w := range words {
		if i == 0 || !small[w] {
			words[i] = strings.ToUpper(string(w[0])) + w[1:]
		}
	}
	return strings.Join(words, " ")
}

func snakeCase(text string) string {
	var r []rune
	prevUnderscore := false
	for _, c := range strings.ToLower(text) {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			r = append(r, c)
			prevUnderscore = false
		} else if unicode.IsSpace(c) {
			if !prevUnderscore {
				r = append(r, '_')
				prevUnderscore = true
			}
		}
	}
	return strings.Trim(string(r), "_")
}

func reverseWords(text string) string {
	var result strings.Builder
	start := 0
	for i, c := range text {
		if unicode.IsSpace(c) {
			word := text[start:i]
			result.WriteString(reverseString(word))
			result.WriteRune(c)
			start = i + 1
		}
	}
	result.WriteString(reverseString(text[start:]))
	return result.String()
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	info := `SENTINEL STRING TRANSFORMER — WELCOME
Available commands:
upper
lower
cap
title 
snake 
reverse
exit`

	fmt.Println(info)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("What are you interested in converting now: ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if strings.TrimSpace(input) == "" {
			continue
		}

		parts := strings.Fields(input)
		cmd := strings.ToLower(parts[0])

		if cmd == "exit" {
			fmt.Println("Goodbye.")
			break
		}

		if len(parts) < 2 {
			fmt.Println("Wrong input")
			continue
		}

		text := input[len(parts[0])+1:]
		var result string

		switch cmd {
		case "upper":
			result = toUpper(text)
		case "lower":
			result = toLower(text)
		case "cap":
			result = capitalize(text)
		case "title":
			result = titleCase(text)
		case "snake":
			result = snakeCase(text)
		case "reverse":
			result = reverseWords(text)
		default:
			fmt.Println("Wrong input")
			continue
		}

		fmt.Println(result)
	}
}
