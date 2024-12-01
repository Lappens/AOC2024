package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	arr1 := []int{}
	arr2 := []int{}
	sum := 0

	// Open .txt
	file, err := os.Open("01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read every Line and split
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		//fmt.Printf("%q\n", test)

		//Convert first and second number to int
		firstNumber, err := strconv.Atoi(text[0])
		if err != nil {
			log.Fatal(err)
		}
		secondNumber, err := strconv.Atoi(text[3])
		if err != nil {
			log.Fatal(err)
		}

		// Append numbers to Arrays
		arr1 = append(arr1, firstNumber)
		arr2 = append(arr2, secondNumber)
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
