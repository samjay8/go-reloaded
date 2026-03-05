package main

import (
	"fmt"
	"goreloaded/transformation"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run. <sample.txt> <result.txt>")
		return
	}

	inputfile := os.Args[1]
	outputfile := os.Args[2]

	content, err := os.ReadFile(inputfile)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	strcont := string(content)
	result := transformation.HexConversion(strcont)
	result = transformation.BinConversion(result)
	result = transformation.UpperCase(result)
	result = transformation.LowerCase(result)
	result = transformation.CapCase(result)
	result = transformation.ArticleReplace(result)
	result = transformation.Punctuation(result)
	result = transformation.QuoteReplace(result)

	err = os.WriteFile(outputfile, []byte(result), 0644)

	if err != nil {
		fmt.Println("Something is wrong with the outputfile, check it out", err)
	}

}
