package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run gourl.go <custom_keyword>")
		os.Exit(1)
	}

	customKeyword := os.Args[1]
	pattern := "(https?://)?(www\\.)?([a-zA-Z0-9.-]+\\.)?" + regexp.QuoteMeta(customKeyword) + "\\.[a-z]+(/\\S*)?"

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		match := regexp.MustCompile(pattern).FindString(line)
		if match != "" {
			fmt.Println(match)
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading input:", scanner.Err())
	}
}
