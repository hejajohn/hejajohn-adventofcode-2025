package main

import (
	"bufio"
	"day2/idchecker"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, Day 2!")

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	idRanges := strings.Split(line, ",")
	invalidIds := 0
	invalidIdsSum := 0

	for _, idRange := range idRanges {
		fmt.Printf("Checking range %s\n", idRange)
		ids := strings.Split(idRange, "-")
		count, sum := idchecker.CheckIdRange(ids[0], ids[1])
		invalidIds += count
		invalidIdsSum += sum
	}

	fmt.Printf("Total invalid IDs: %d\n", invalidIds)
	fmt.Printf("Sum of invalid IDs: %d\n", invalidIdsSum)
}
