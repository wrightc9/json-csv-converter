package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func menu() {
	// Print Menu
	fmt.Println("1. Convert JSON to CSV")
	fmt.Println("2. Convert CSV to JSON")
	fmt.Println("3. Exit")

	// Get user input
	fmt.Println("Enter your choice: ")
	var choice int
	fmt.Scanln(&choice)

	// Handle user input
	if choice < 3 && choice > 0 {
		convert(choice)
	} else if choice == 3 {
		fmt.Println("Exiting...")
	} else {
		fmt.Println("Invalid choice")
	}
}

// Function to convert JSON to CSV or vice versa
func convert(convType int) {
	var inFile *os.File
	var outFile *os.File

	// Prompt user for file names
	// 1 = JSON to CSV
	// 2 = CSV to JSON
	if convType == 1 {
		inFile = promptFileNameReq(".json", "o")
		outFile = promptFileNameReq(".csv", "c")

		defer inFile.Close()
	} else {
		inFile = promptFileNameReq(".csv", "o")
		outFile = promptFileNameReq(".json", "c")

		defer inFile.Close()
	}

	// Read file line by line
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		if convType == 1 {
			// Convert JSON to CSV
			outFile.WriteString(jsonToCSV(line))
		} else {
			// Convert CSV to JSON
			outFile.WriteString(csvToJSON(line))
		}
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}

func jsonToCSV(line string) string {
	// convert json line to csv line
	return line
}

func csvToJSON(line string) string {
	// convert csv line to json line
	return line
}

func promptFileNameReq(ext string, flag string) *os.File {
	var fileName string
	var file *os.File
	for {
		fmt.Printf("Enter %s file name: ", ext)
		fmt.Scanln(&fileName)

		if fileName == "" {
			fmt.Println("Closing...")
			os.Exit(0)
		}
		// Check if file name ends with .json
		if !strings.HasSuffix(fileName, ext) {
			fmt.Printf("Invalid file name. File must end with %s extension.\n", ext)
			continue
		}

		var err error
		if flag == "o" {
			file, err = os.Open(fileName)
			if err != nil {
				fmt.Println("Error opening file:", err)
				continue
			}
		} else {
			file, err = os.Create(fileName)
			if err != nil {
				fmt.Println("Error creating file:", err)
				continue
			}
		}

		break
	}

	return file
}
