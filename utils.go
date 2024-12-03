package utils

import (
	"log"
	"os"
	"strconv"
)

func ReadFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func StringArrayToInt(array []string) []int {
	arrInt := make([]int, len(array))
	for index, value := range array {
		i, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		arrInt[index] = i
	}
	return arrInt
}
