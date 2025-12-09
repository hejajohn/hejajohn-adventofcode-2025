package main

import (
	"bufio"
	"day3/battery"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	joltSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		cells := battery.StringToCellArray(line)
		jolt := battery.GetMaxJolt(cells)
		joltSum += jolt

		fmt.Printf("Max jolt: %d\n", jolt)
	}
	fmt.Printf("Jolt sum: %d\n", joltSum)
}
