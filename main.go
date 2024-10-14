package main

import (
	"flag"
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
	var indentationString string
	var depth int
	var startBy int

	flag.StringVar(&indentationString, "indent", "  ", "Indentation string ('  ','tab')")
	flag.IntVar(
		&depth,
		"depth",
		3,
		"Depth of the table of contents (1,2,...,6). 2 will print only H1 and H2",
	)
	flag.IntVar(&startBy, "start-by", 1, "Start by specific header (1,2,...,6). 2 will start by H2")

	flag.Parse()

	var filePath string
	if len(flag.Args()) < 1 {
		fmt.Println("Usage: markdown-table-of-contents [markdownfile.md]")
		fmt.Println("Type 'markdown-table-of-contents -h' for help")
		os.Exit(1)
	}
	filePath = flag.Args()[0]

	md := markdown.New(filePath)
	err := md.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var strs []string = []string{}
	for _, section := range md.Sections {
		if section.SectionType == markdown.NullSection {
			continue
		}

		numberOfSpaces := len(string(section.SectionType)) - 1
		if numberOfSpaces+1 > depth {
			continue
		}
		if numberOfSpaces+1 < startBy {
			continue
		}
		var str string
		str += strings.Repeat(indentationString, numberOfSpaces)
		str += "- [" + section.Text + "](#" + formatTitle(section.Text) + ")"
		strs = append(strs, str)
	}

	// If all start with indent, remove all
	for {
		allStartWithIndent := true
		for _, str := range strs {
			if !strings.HasPrefix(str, indentationString) {
				allStartWithIndent = false
				break
			}
		}
		if allStartWithIndent {
			for i := range strs {
				strs[i] = strs[i][len(indentationString):]
			}
		} else {
			break
		}
	}

	for _, str := range strs {
		fmt.Println(str)
	}
}
