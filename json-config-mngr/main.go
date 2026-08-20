package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Address struct {
	City    string `json:"city"`
	Pincode int    `json:"pincode"`
}

type Config struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

func loadConfig() (map[string]any, error) {
	data, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return nil, err
	}

	var config map[string]any
	err = json.Unmarshal(data, &config)

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}
	return config, nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("error loading configuration: %v\n", err)
	}

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	inputWords := strings.Fields(input)

	if len(inputWords) < 2 {
		return
	}

	if inputWords[0] != "get" {
		fmt.Printf("use prefix 'get' instead of '%v'\n", inputWords[0])
		return
	}
	jsonKey := inputWords[1]
	fmt.Printf("json key %v\n", jsonKey)

	// bring value from json using jsonKey
	jsonValue, ok := config[jsonKey]
	if ok {
		fmt.Printf("json vlaue %v\n", jsonValue)
	} else {
		fmt.Printf("no '%v' key found in json\n", jsonKey)
		return
	}

}
