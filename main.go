package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"crea/parkinglot"
)

func main() {
	if len(os.Args) > 1 {
		filename := os.Args[1]
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error opening the file:", err)
			return
		}
		defer file.Close()
		processFile(file)
	} else {
		fmt.Println("Enter commands (type 'exit' to quit):")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.ToLower(line) == "exit" {
				break
			}
			parkinglot.ProcessCommand(line)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from input:", err)
		}
	}
}

func processFile(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parkinglot.ProcessCommand(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
}