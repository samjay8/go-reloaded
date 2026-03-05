package transformation

import "regexp" // Import reguaar expression package

func Punctuation(text string) string {
	word1 := regexp.MustCompile(`([\s]+)([.,!?:;]+)`) // Compiles pattern tht can be found in the sentence
	result := word1.ReplaceAllString(text, "$2")      // Replaces the pattern to remove spaces

	word2 := regexp.MustCompile(`([.,!?:;]+)([a-zA-z0-9])`) // Complies pattern to note one or more punctuation marks and only letter and numbers
	result = word2.ReplaceAllString(result, "$1 $2")        // Replace it with the punctuation marks, a space, and a letter or number
	return result
}
