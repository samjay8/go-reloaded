package transformation

import "regexp"

func QuoteReplace(text string) string {
	word := regexp.MustCompile(`'\s+(.*?)\s+'`)   // Checks for the pattern in the sentence
	result := word.ReplaceAllString(text, "'$1'") // Replace the pattern to remove spaces in the content.

	return result
}
