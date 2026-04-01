package main

import (
	"bufio"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	f1, err := os.Open("sample.txt")
	check(err)
	defer f1.Close()

	f2, err := os.Create("result.txt")
	check(err)
	defer f2.Close()

	text := bufio.NewScanner(f1)
	writer := bufio.NewWriter(f2)

	for text.Scan() {
		input := text.Text()

		text := input

		text = toUppertoLower(text)
		text = BinToDecimal(text)
		text = Allhex(text)
		text = Quote(text)
		text = articleA(text)
		text = fixPunctuation(text)
		text = capitalize(text)

		_, err := writer.WriteString(text + "\n")
		if err != nil {
			log.Fatal(err)
		}
		writer.Flush()
	}
	if err := text.Err(); err != nil {
		log.Fatal(err)
	}
}
