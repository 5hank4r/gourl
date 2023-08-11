package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: urlexgo <custom_keyword>")
		os.Exit(1)
	}

	customKeyword := os.Args[1]
	pattern := fmt.Sprintf(`(https?://)?(www\.)?([a-zA-Z0-9.-]+\.)?%s\.[a-z]+(/\S*)?`, regexp.QuoteMeta(customKeyword))
	compiledPattern := regexp.MustCompile(pattern)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if compiledPattern.MatchString(line) {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
		os.Exit(1)
	}
}
