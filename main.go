package main

import (
	"github.com/joho/godotenv"
    "log"
    "os"
	"strings"
)

func main() {
  	file := ".env"
	err := godotenv.Load(file)
	EnvMap := make(map[string]string)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	for _, e := range os.Environ() {
		subStrings := strings.Split(e, "=")
		key := subStrings[0]
		value := subStrings[1]
		EnvMap[key] = value
	}
	
	for key, value := range EnvMap {
		log.Println(key, value)
	}
}