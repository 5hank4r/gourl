package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gourl <custom_keyword>")
		os.Exit(1)
	}

	customPattern := os.Args[1]
	pattern := fmt.Sprintf(`(https?://)?(www\.)?%s[^ /]+`, regexp.QuoteMeta(customPattern))

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if matched, _ := regexp.MatchString(pattern, line); matched {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
		os.Exit(1)
	}
}
