package transformation

import "strings"

func ArticleReplace(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ { // Iterates through the entire slice of words from the first character
		if words[i] == "a" || words[i] == "A" { //If word is a or A
			if i+1 < len(words) {
				switch words[i+1][0] { // Checks if the first character of the next word is a vowel or h
				case 'a', 'e', 'i', 'o', 'u', 'h', 'A', 'E', 'I', 'O', 'U', 'H':
					if words[i] == "a" { // if the vowel is 'a', replace with 'an
						words[i] = "an"
					} else {
						words[i] = "An" // if the vowel is 'A', replace with 'An'
					}
				}
			}
		}
	}
	return strings.Join(words, " ") // join the slice of strings together and return it back as a string.
}
