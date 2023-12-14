// Solution to Write Your Own wc Tool colding challenge
// 2023-12-13 - MFSH

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	usage = `
	Usage: ccwc [OPTION] [FILE]

	-c print the byte counts
	-m print the character counts
	-l print the newlines counts 
	-w print the word counts
	`
)

func main() {
	var filename string
	var option string

	// Check if an option is provided
	if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-") {
		option = os.Args[1]

		// Check if the provided option is valid
		validOptions := map[string]bool{"-c": true, "-l": true, "-w": true, "-m": true}
		if !validOptions[option] {
			fmt.Println("Invalid option:", option)
			fmt.Println("Valid options are: -c, -l, -w, -m")
			return
		}

		// Check if a filename is provided after the option
		if len(os.Args) > 2 {
			filename = os.Args[2]
		}
	} else if len(os.Args) > 1 {
		// if no option provided, use the filename from os.Args[1]
		filename = os.Args[1]
	} else {
		// if no command-line arguments given, input is from a pipe
		filename = "stdin"
	}

	var lineCount, byteCount, wordCount, charCount int

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Input is from a pipe, process it
		_ = stat
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			lineCount++
			byteCount += len(scanner.Bytes())
			wordCount += countWords(scanner.Text())
			charCount += utf8.RuneCountInString(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from pipe:", err)
			return
		}
	} else {
		// Input is not from a pipe, open the file
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		// Iterate through each line in the file
		for scanner.Scan() {
			lineCount++
			byteCount += len(scanner.Bytes())
			wordCount += countWords(scanner.Text())
			charCount += utf8.RuneCountInString(scanner.Text())
		}

		// Check for scanner errors
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	}

	switch option {
	case "-c":
		fmt.Println(byteCount, filename)
	case "-l":
		fmt.Println(lineCount, filename)
	case "-w":
		fmt.Println(wordCount, filename)
	case "-m":
		fmt.Println(charCount, filename)
	default:
		fmt.Println(lineCount, wordCount, byteCount, filename)
	}
}

// countWords counts the number of words in a given text
func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}
