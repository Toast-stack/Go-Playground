package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Open the log file
	file, err := os.Open("sample.log")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Define a regex pattern to find failed logins
	pattern := `ERROR Login failed: IP=([\d.]+)`
	re := regexp.MustCompile(pattern)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Match the regex pattern
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			// Extract and print the IP address
			fmt.Println("Failed login attempt detected from IP:", matches[1])
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
