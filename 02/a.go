package main

import (
	utils "aoc2024"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	safeCalc := 0
	// Open .txt
	file := utils.ReadFile("02/input.txt")
	defer file.Close()
	// Read every Line and split
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Fields(scanner.Text())
		//Convert Strings to Int
		convInt := utils.StringArrayToInt(text)
		if ascOrDesc(convInt) {
			safeCalc++
		}
	}
	fmt.Println(safeCalc)
}

func ascOrDesc(array []int) bool {
	// Ascending
	if array[0] < array[1] {
		for i := 0; i < len(array)-1; i++ {
			if (array[i+1]-array[i] > 3) || (array[i]-array[i+1] == 0) {
				return false
			}
			if array[i] > array[i+1] {
				return false
			}
		}
		// Descending
	} else {
		for i := 0; i < len(array)-1; i++ {
			if (array[i]-array[i+1] > 3) || (array[i]-array[i+1] == 0) {
				return false
			}
			if array[i] < array[i+1] {
				return false
			}
		}
	}
	return true
}
