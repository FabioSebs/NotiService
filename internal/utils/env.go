package utils

import (
	"log"
	"strconv"
)

func GetInt(key string) int {
	number, err := strconv.Atoi(key)
	if err != nil {
		log.Fatal("invalid env")
	}
	return number
}
