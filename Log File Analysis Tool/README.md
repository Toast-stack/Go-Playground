# Project Documentation: Log File Analysis Tool

## Overview
The Log File Analysis Tool is a command-line application written in Go that processes log files to identify specific patterns, such as failed login attempts. This project demonstrates proficiency in file handling, pattern matching, and basic cybersecurity analysis.

## Key Features
1. **Log Parsing**: Reads and analyzes log files line by line.
2. **Regex Matching**: Detects failed login attempts using regular expressions.
3. **Pattern Detection**: Extracts and displays suspicious IP addresses associated with failed logins.

## Code Example
Here’s a snippet of the program that highlights its core functionality:
``` Go
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
```
## How It Works
1. The program opens a log file (`sample.log`) and reads it line by line.
2. It uses a regular expression (`regex`) to identify lines containing failed login attempts (e.g., entries marked `ERROR`).
3. The regex extracts the IP address from the matching lines.
4. The program outputs the extracted IP addresses. helping to detect suspicious activity.

## Why This Project Is Important
1. **Hands-On Experience**: Demonstrates file I/O operations, regex usage, and real-world log analysis, all of which are essential cybersecurity skills.
2. **Threat Detection**: Provides a simple yet effective way to identify failed login attempts, a common indicator of brute-force attacks or unauthorized access attempts.
3. **Scalable Design**: Serves as a foundation for more advanced log analysis tools with features like severity classification or automated reporting.
4. **Portfolio Value**: Highlights problem-solving and practical coding skills, making it a standout project for technical roles.

## References for Further Learning
To support understanding of the concepts used in this project, here are some helpful resources:
- [Go's os Package Documentation](https://pkg.go.dev/os): Learn about file handling in Go.
- [Go's bufio Package Documentation](https://pkg.go.dev/bufio): Understand buffered file reading.
- [Regular Expressions in Go](https://pkg.go.dev/regexp): A guide to using regex for pattern matching in Go.

## Potential Enhancements
1. **Count IP Occurrences**: Track how often each IP appears in the log file to identify repeated login failures.
2. **Severity Categorization**: Extend regex patterns to analyze different log levels (e.g., INFO, WARNING, ERROR).
3. **Report Generation**: Export analysis results to a new file (e.g., `report.log`).
4. **Real-Time Monitoring**: Adapt the tool to monitor logs in real-time for immediate alerts.
