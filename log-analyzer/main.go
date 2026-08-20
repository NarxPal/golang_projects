package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("app.log")
	if err != nil {
		fmt.Printf("eror reading file: %v\n", err)
	}

	defer file.Close() // dont' close immediately

	// read input in tokens (here, token would contain a single line)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text() // returns newly generated token( which is a line )
		fmt.Println(line)      // print the scanned line
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}
