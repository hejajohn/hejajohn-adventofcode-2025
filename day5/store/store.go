package store

import (
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	From int
	To   int
}

func (r1 *Range) Overlaps(r2 Range) bool {
	return r1.From <= r2.To && r2.From <= r1.To
}

func (r1 *Range) Expand(r2 Range) {
	(*r1).From = min(r1.From, r2.From)
	(*r1).To = max(r1.To, r2.To)
}

func (r *Range) ContainsId(id int) bool {
	return id >= r.From && id <= r.To
}

func (r *Range) ToString() string {
	return fmt.Sprintf("(%d-%d)", r.From, r.To)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetOverlappingRange(ranges []Range, r Range) *Range {
	for i := range ranges {
		if r == ranges[i] {
			continue
		}
		if r.Overlaps(ranges[i]) {
			return &ranges[i]
		}
	}
	return nil
}

func CreateRangeFromString(input string) Range {
	split := strings.Split(input, "-")

	if len(split) != 2 {
		panic("Unexpected range")
	}

	from, errFrom := strconv.Atoi(split[0])
	to, errTo := strconv.Atoi(split[1])

	if errFrom != nil || errTo != nil {
		panic("Invalid range values")
	}

	return Range{From: from, To: to}
}

func CreateIdFromString(input string) int {
	id, err := strconv.Atoi(input)

	if err != nil {
		panic("Invalid id")
	}

	return id
}
