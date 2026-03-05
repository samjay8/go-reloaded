package transformation

import (
	"fmt"
	"strconv"
	"strings"
)

func HexConversion(text string) string {
	words := strings.Fields(text) // Split the text into a slice of words
	for i := 1; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			hexdec, err := strconv.ParseInt(words[i-1], 16, 64) // convert the string into decimal
			if err != nil {
				fmt.Println("Error", err) // handle error
			} else {
				words[i-1] = fmt.Sprintf("%d", hexdec) //Convert the decimal integer back to a string and replace the original word
			}
			words = append(words[:i], words[i+1:]...) //this removes the marker (hex)
			i--                                       //reverse i to note the removed element so no word is skipped
		}
	}
	return strings.Join(words, " ") //return the new slice(words) as a string
}

func BinConversion(text string) string {
	words := strings.Fields(text) // Change the text into a slice of strings
	for i := 1; i < len(words); i++ {
		if words[i] == "(bin)" && i > 0 {
			bindec, err := strconv.ParseInt(words[i-1], 2, 64) // convert the string into decimal
			if err != nil {
				fmt.Println("Error", err) //handle error
			} else {
				words[i-1] = fmt.Sprintf("%d", bindec) //Convert the decimal integer back to a string and replace the original word

			}
			words = append(words[:i], words[i+1:]...) //this removes the marker (bin)
			i--                                       //reverse i to note the removed element so no word is skipped
		}
	}
	return strings.Join(words, " ") //return the new slice(words) as a string
}
