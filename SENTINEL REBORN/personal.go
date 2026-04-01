

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

func main() {
	fmt.Println(fixPunctuation(("Punctuation tests are    ... kinda boring  , do you think?")))
}
