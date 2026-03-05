package transformation

import (
	"fmt"
	"strconv"
	"strings"
)

func UpperCase(text string) string {
	words := strings.Fields(text) // Split the text into a slice of words

	for i := 1; i < len(words); i++ {
		if words[i] == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1]) // Convert the word to Uppercase

			words = append(words[:i], words[i+1:]...) // Remove the marker from the sentence
			i--                                       // Reverse the string to note the removal of the marker
		}

		if words[i] == "(up," && i > 0 {

			value := strings.TrimSuffix(words[i+1], ")") // Trim away the suffix not needed from the string
			number, err := strconv.Atoi(value)           // convert the string to int
			if err != nil {
				fmt.Println("Error", err) //handle error
			}
			for a := i - 1; a >= i-number && a >= 0; a-- { // A mini loop to iterate over hnumber of word to convert
				words[a] = strings.ToUpper(words[a]) // COnvert the words to Uppercase
			}

			words = append(words[:i], words[i+2:]...) // Remove the markers from the sentence
			i--                                       // Reverse the string to note the removal of the marker
		}
	}
	return strings.Join(words, " ") // Convert the slice of words back to string and return it
}

func LowerCase(text string) string {
	words := strings.Fields(text) // Split the text into a slice of words

	for i := 1; i < len(words); i++ {
		if words[i] == "(low)" && i > 0 {
			words[i-1] = strings.ToLower(words[i-1]) // Convert the word to Lowrecase

			words = append(words[:i], words[i+1:]...) // Remove the marker from the sentence
			i--                                       // Reverse the string to note the removal of the marker
		}

		if words[i] == "(low," && i > 0 { //

			value := strings.TrimSuffix(words[i+1], ")") // Trim away the suffix not needed from the string
			number, err := strconv.Atoi(value)           // convert the string to int
			if err != nil {
				fmt.Println("Error", err) //handle error
			}
			for a := i - 1; a >= i-number && a >= 0; a-- { // A mini loop to iterate over hnumber of word to convert
				words[a] = strings.ToLower(words[a]) // COnvert the words to Lowercase
			}

			words = append(words[:i], words[i+2:]...) // Remove the markers from the sentence
			i--                                       // Reverse the string to note the removal of the marker
		}
	}
	return strings.Join(words, " ") // Convert the slice of words back to string and return it
}

func CapCase(text string) string {
	words := strings.Fields(text)

	for i := 1; i < len(words); i++ {
		if words[i] == "(cap)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1][:1]) + strings.ToLower(words[i-1][1:]) // Capitalize the first word
			words = append(words[:i], words[i+1:]...)                                      // Remove the marker
			i--
		}
		if words[i] == "(cap," && i > 0 { // check for multiple cases
			value := strings.TrimSuffix(words[i+1], ")") // Trim away the suffix not needed from the string
			number, err := strconv.Atoi(value)

			if err != nil {
				fmt.Println("Error", err)
			}
			for a := i - 1; a >= i-number && a >= 0; a-- {
				words[a] = strings.ToUpper(words[a][:1]) + strings.ToLower(words[a][1:])
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}
