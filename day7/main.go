package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var beams []int
	var splitters [][]int
	var totalSplits int

	for scanner.Scan() {
		s := make([]int, 0)
		for x, b := range scanner.Bytes() {
			switch b {
			case 'S':
				beams = append(beams, x)
			case '^':
				s = append(s, x)
			}
		}
		splitters = append(splitters, s)
	}

	for y, row := range splitters {
		if y == 0 || y%2 == 1 {
			continue
		}

		for _, s := range row {
			for _, b := range beams {
				if s == b {
					totalSplits++

					bl := b - 1
					if !slices.Contains(row, bl) && !slices.Contains(beams, bl) {
						beams = append(beams, bl)
					}
					br := b + 1
					if !slices.Contains(row, br) && !slices.Contains(beams, br) {
						beams = append(beams, br)
					}

					for i, beam := range beams {
						if b == beam {
							beams = append(beams[:i], beams[i+1:]...)
						}
					}
				}
			}
		}
	}

	fmt.Println(beams)
	fmt.Printf("Total splits: %d", totalSplits)
}
