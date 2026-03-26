package main

import (
	"fmt"
	"log"
	"os"
)

const FILENAME = "data.txt"

func main() {
	fmt.Println("Welcome to FileReader Writer!")



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