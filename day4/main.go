package main

import (
	"bufio"
	"day4/matrix"
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

	var mat [][]bool

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		mat = append(mat, matrix.StringToRow(line))
	}

	tpCount := 0

	for p := range matrix.All(mat) {
		if !mat[p.Y][p.X] {
			continue
		}

		// fmt.Printf("Check neighbors (%d, %d)\n", p.X, p.Y)
		neighbors := 0
		for _, val := range matrix.Neighbors(mat, p) {
			if val {
				// fmt.Printf("(%d, %d)\n", n.X, n.Y)
				neighbors++
			}
		}
		if neighbors < 4 {
			tpCount++
		}
	}

	fmt.Printf("TP rolls: %d\n", tpCount)
}
