package main

import (
	"bufio"
	"day6/equation"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	rNum, err := regexp.Compile(`^\s{0,}(\d+)\s{0,}`)
	if err != nil {
		panic(err)
	}
	rOp, err := regexp.Compile(`^([+\-*/]+)\s{0,}`)
	if err != nil {
		panic(err)
	}

	var rows []string

	for scanner.Scan() {
		line := scanner.Text()

		rows = append(rows, line)
	}

	var m []string
	var totalSum int

	for {
		var numbers []int
		var operator string

		for i, line := range rows {
			if i < len(rows)-1 {
				m = rNum.FindStringSubmatch(line)
				if m == nil {
					panic("Unexpected number input")
				}
				n, _ := strconv.Atoi(m[1])
				numbers = append(numbers, n)
			} else {
				m = rOp.FindStringSubmatch(line)
				if m == nil {
					panic("Unexpected operation input")
				}
				operator = m[1]
			}
			rows[i] = rows[i][len(m[0]):]
		}

		eq := equation.CreateEquation(numbers, operator)
		// fmt.Printf("%s = %d\n", eq.ToString(), eq.Compute())
		totalSum += eq.Compute()

		if rows[0] == "" {
			break
		}
	}

	fmt.Printf("Total sum: %d\n", totalSum)
}
