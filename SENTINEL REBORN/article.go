package main

import (
	"strings"
)

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
