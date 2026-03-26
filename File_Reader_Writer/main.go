package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const FILENAME = "data.txt"

func main() {
	fmt.Println("Welcome to FileReader Writer!")

	// Reading Initial Data
	initialContent := ReadFile()
	fmt.Println("Initial Content: ")
	fmt.Println(initialContent)

	fmt.Println("------------")

	// Write Data

	fmt.Println("Enter a text to write to a file: ")
	scanner := bufio.NewScanner(os.Stdin)

	var allLines []string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" && len(allLines) > 0 {
			break
		}

		allLines = append(allLines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input...")
		return
	}

	WriteFile(allLines)

	// Reading Afterward data

	fmt.Println("-----------")

	endContent := ReadFile()
	fmt.Println("End Content: ")
	fmt.Println(endContent)

}

func WriteFile(lines []string){
	file, err := os.Create(FILENAME)

	if err != nil {
		log.Fatalln("Error creating file!")
	}

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			log.Fatalln("Error writing file...")
		}
	}
}

func ReadFile() string{
	content, err := os.ReadFile(FILENAME)
	if err != nil {
		log.Fatalln("Error reading file...")
	}
	return string(content)
}