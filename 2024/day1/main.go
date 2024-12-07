package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("err opening file %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var location1, location2 []int
	for scanner.Scan() {
		str := strings.Fields(scanner.Text())

		num1, err := strconv.Atoi(str[0])
		if err != nil {
			panic("error coy")
		}
		num2, err := strconv.Atoi(str[1])
		if err != nil {
			panic("error coy")
		}

		location1 = append(location1, num1)
		location2 = append(location2, num2)

	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("err reading file %v", err)
	}

	total := 0
	numOccurance := make(map[int]int)
	for _, item := range location1 {
		val, ok := numOccurance[item]
		if !ok {
			count := searchOcurance(item, location2)
			numOccurance[item] = count
			val = count
		}
		total += val * item
	}
	fmt.Println(total)
	fmt.Printf(
		"total time taken %v ms(%v microseconds)",
		time.Since(startTime).Milliseconds(),
		time.Since(startTime).Microseconds(),
	)
}

func searchOcurance(str int, slice []int) int {
	count := 0
	for _, item := range slice {
		if item == str {
			count++
		}
	}
	return count
}
