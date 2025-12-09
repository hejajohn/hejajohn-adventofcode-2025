package battery

import (
	"cmp"
	"fmt"
	"slices"
)

const zero byte = '0'

type Cell struct {
	value int
	pos   int
}

func StringToCellArray(input string) []Cell {
	chars := []byte(input)
	cells := make([]Cell, len(chars))
	for i, char := range chars {
		cells[i] = Cell{
			value: int(char - zero),
			pos:   i,
		}
	}
	return cells
}

func GetMaxJolt(cells []Cell) int {
	sortedCells := getSortedCellArray(cells)
	// printCells(sortedCells)
	return getMaxJoltRecursive(0, sortedCells)
}

func getSortedCellArray(cells []Cell) []Cell {
	sortedCells := make([]Cell, len(cells))
	copy(sortedCells, cells)
	slices.SortFunc(sortedCells, func(a, b Cell) int {
		return -cmp.Compare(a.value, b.value)
	})
	return sortedCells
}

func getMaxJoltRecursive(index int, cells []Cell) int {
	if index >= len(cells)-1 {
		return 0
	}
	cur := cells[index]
	jolt := 0
	for _, cell := range cells {
		if cell.pos > cur.pos {
			jolt = getJolt(cur.value, cell.value)
			break
		}
	}
	// fmt.Printf("cur: %d rest: ", cur.value)
	// printCells(rest)
	// fmt.Printf("jolt: %d next max jolt: %d\n", jolt, getJolt(cells[index+1].value, 9))
	if index < len(cells)-1 && jolt >= getJolt(cells[index+1].value, 9) {
		return jolt
	}
	return max(jolt, getMaxJoltRecursive(index+1, cells))
}

func getJolt(a, b int) int {
	return a*10 + b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printCells(cells []Cell) {
	for _, cell := range cells {
		fmt.Printf("%d(%d) ", cell.value, cell.pos)
	}
	fmt.Println()
}
