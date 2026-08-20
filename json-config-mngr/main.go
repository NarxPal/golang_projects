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
	var jsonValue any
	var nestedMap map[string]any
	var ok bool

	if strings.Contains(jsonKey, ".") {
		parts := strings.Split(jsonKey, ".")
		fmt.Printf("parts : %v\n", parts)
		jsonPart0Val, exists := config[parts[0]]
		if !exists {
			fmt.Printf("key %s not found\n", parts[0])
			return
		}

		nestedMap, ok = jsonPart0Val.(map[string]any)
		if !ok {
			fmt.Printf("%s is not a nested object\n", parts[0])
			return
		}

		jsonValue, exists = nestedMap[parts[1]]
		if !exists {
			fmt.Printf("key %s not found inside %s\n", parts[1], parts[0])
			return
		}

		fmt.Printf("json value: %v\n", jsonValue)

	} else {
		jsonValue, ok = config[jsonKey]

		fmt.Printf("json value: %v\n", jsonValue)
	}

}
