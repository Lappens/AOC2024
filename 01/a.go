package main

import (
	utils "aoc2024"
	"bufio"
	"fmt"
	"slices"
	"strings"
)

func main() {
	arr1 := []int{}
	arr2 := []int{}
	sum := 0

	// Open .txt
	file := utils.ReadFile("01/input.txt")

	// Read every Line and split
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		//fmt.Printf("%q\n", test)

		//Convert first and second number to int
		textToNumbers := utils.StringArrayToInt(text[0])

		// Append numbers to Arrays
		arr1 = append(arr1, textToNumbers[0])
		arr2 = append(arr2, textToNumbers[1])
	}
	// fmt.Println(arr1)

	// Sort Array min -> max
	slices.Sort(arr1)
	slices.Sort(arr2)

	// fmt.Println(arr1)
	// Loop Arrays and calculate
	for i := 0; i < len(arr1); i++ {
		var difference = arr1[i] - arr2[i]
		if difference < 0 {
			difference = -difference
		}
		sum += difference
	}

	fmt.Println(sum)
}
