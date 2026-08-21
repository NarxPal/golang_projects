package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var INFO int = 0
	var WARN int = 0
	var ERROR int = 0
	var errLog map[string]int
	errMsgs := []string{}

	file, err := os.Open("app.log")
	if err != nil {
		fmt.Printf("eror reading file: %v\n", err)
	}

	defer file.Close() // dont' close immediately

	// read input in tokens (here, token would contain a single line)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()        // returns newly generated token( which is a line )
		fmt.Println(line)             // print the scanned line
		words := strings.Fields(line) //
		for i, word := range words {
			if word == "INFO" || word == "WARNING" || word == "ERROR" {
				fmt.Printf("%s found\n", word) // print the level of log for a line(token)
			}

			switch word {
			case "INFO":
				INFO++

			case "WARNING":
				WARN++

			case "ERROR":
				ERROR++

				errLog = make(map[string]int)

				// taking words present after "ERROR" in line
				errorWords := words[i+1:]

				// combine those words to form complete string of error statement
				errorMsg := strings.Join(errorWords, " ")

				errMsgs = append(errMsgs, errorMsg)

				for _, key := range errMsgs {
					errLog[key]++
				}

			}

		}
	}

	fmt.Println("INFO :", INFO)
	fmt.Println("WARN :", WARN)
	fmt.Println("ERROR :", ERROR)
	fmt.Println("total count:", INFO+WARN+ERROR)

	fmt.Println("Err msgs:", errLog)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}
