// Package files provides funcs to files
package files

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// ContentFileByLines :
// return lines of file in a slice
func ContentFileByLines(file string) ([]string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
}

// CommentIfFind :
// comment lines with regex, return lines commented
func CommentIfFind(file, regex, commentPattern string) ([]string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	commentedLines := make([]string, 0)

	for i, line := range lines {
		matched, errInLoop := regexp.MatchString(regex, line)
		if errInLoop != nil {
			return nil, errInLoop
		}

		if matched {
			lines[i] = commentPattern + line
			fmt.Printf("File : %s\nFound at line %d : %s\n\n", file, i+1, line)
			commentedLines = append(commentedLines, line)
		}
	}
	return commentedLines, nil
}
