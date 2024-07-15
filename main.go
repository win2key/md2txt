package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/russross/blackfriday/v2"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: md2txt <input.md> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	mdData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	// Convert Markdown to HTML
	htmlData := blackfriday.Run(mdData)

	// Remove HTML tags to get plain text
	plainText := stripHTMLTags(string(htmlData))

	err = ioutil.WriteFile(outputFile, []byte(plainText), 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		return
	}

	fmt.Printf("Converted %s to %s successfully.\n", inputFile, outputFile)
}

// stripHTMLTags removes HTML tags from a string
func stripHTMLTags(html string) string {
	var sb strings.Builder
	inTag := false

	for _, r := range html {
		switch {
		case r == '<':
			inTag = true
		case r == '>':
			inTag = false
		case !inTag:
			sb.WriteRune(r)
		}
	}

	return sb.String()
}
