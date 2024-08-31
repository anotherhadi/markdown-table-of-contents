package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/anotherhadi/markdown"
)

func formatTitle(title string) string {
	str := strings.TrimSpace(title)
	str = strings.ReplaceAll(str, " ", "-")
	str = regexp.MustCompile(`[^a-zA-Z0-9 _-]+`).ReplaceAllString(str, "")
	str = strings.ToLower(str)
	return str
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: markdown-table-of-contents [markdownfile.md] <Identation string>")
		os.Exit(1)
	}

	indentationString := "  " // By default, use 2 spaces
	if len(os.Args) > 2 {
		indentationString = os.Args[2]
	}

	filePath := os.Args[1]
	md := markdown.New(filePath)
	err := md.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	str := ""
	for _, section := range md.Sections {
		if section.SectionType == markdown.NullSection {
			continue
		}

		numberOfSpaces := len(string(section.SectionType)) - 1
		str += strings.Repeat(indentationString, numberOfSpaces)
		str += "- [" + section.Text + "](#" + formatTitle(section.Text) + ")"
		str += "\n"
	}
	fmt.Println(str)
}
