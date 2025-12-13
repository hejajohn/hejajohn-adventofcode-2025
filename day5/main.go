package main

import (
	"bufio"
	"day5/store"
	"fmt"
	"os"
)

// Try 1: 403, too low.

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ranges []store.Range

	scanningRanges := true
	freshIdCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if scanningRanges {
				scanningRanges = false
				fmt.Printf("Read %d ranges\n", len(ranges))
				for _, r := range ranges {
					fmt.Println(r.ToString())
				}
				fmt.Println()
				continue
			} else {
				break
			}
		}
		if scanningRanges {
			var r1 store.Range = store.CreateRangeFromString(line)
			var r2 *store.Range = store.GetOverlappingRange(ranges, r1)

			if r2 != nil {
				for {
					r2.Expand(r1)

					for i := range ranges {
						if ranges[i] == r1 {
							ranges = append(ranges[:i], ranges[i+1:]...)
							break
						}
					}

					// Continue trying to merge with any other overlapping ranges
					r1 = *r2
					r2 = store.GetOverlappingRange(ranges, r1)
					if r2 == nil {
						break
					}
				}
			} else {
				ranges = append(ranges, r1)
			}
		} else {
			id := store.CreateIdFromString(line)
			for _, r := range ranges {
				if r.ContainsId(id) {
					freshIdCount++
					break
				}
			}
		}
	}

	fmt.Printf("Found %d fresh Id:s\n", freshIdCount)
}
