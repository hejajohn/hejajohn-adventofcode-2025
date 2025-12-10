package matrix

import "iter"

type Pos struct {
	X int
	Y int
}

func StringToRow(input string) []bool {
	chars := []byte(input)
	cells := make([]bool, len(chars))
	for i, char := range chars {
		cells[i] = char == '@'
	}
	return cells
}

func Neighbors(mat [][]bool, pos Pos) iter.Seq2[Pos, bool] {
	return func(yield func(Pos, bool) bool) {
		p := Pos{}

		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}

				p.X = pos.X + j
				p.Y = pos.Y + i

				if p.X < 0 || p.X >= len(mat[0]) || p.Y < 0 || p.Y >= len(mat) {
					continue
				}

				if !yield(p, mat[p.Y][p.X]) {
					return
				}
			}
		}
	}
}

func All(mat [][]bool) iter.Seq[Pos] {
	return func(yield func(Pos) bool) {
		p := Pos{}
		for i := 0; i < len(mat); i++ {
			for j := 0; j < len(mat[0]); j++ {
				p.X = j
				p.Y = i
				if !yield(p) {
					return
				}
			}
		}
	}
}
