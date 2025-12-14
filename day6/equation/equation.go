package equation

import (
	"fmt"
	"strconv"
	"strings"
)

type Equation struct {
	numbers  []int
	operator string
}

func CreateEquation(numbers []int, operator string) Equation {
	if operator != "+" && operator != "*" {
		panic("Invalid operator")
	}
	return Equation{numbers, operator}
}

func (e *Equation) ToString() string {
	var numStrings []string
	for _, n := range e.numbers {
		numStrings = append(numStrings, strconv.FormatInt(int64(n), 10))
	}
	return strings.Join(numStrings, fmt.Sprintf(" %s ", e.operator))
}

func (e *Equation) Compute() int {
	switch e.operator {
	case "+":
		return e.computeAddition()
	case "*":
		return e.computeMultiplication()
	}
	panic("Invalid operator")
}

func (e *Equation) computeAddition() int {
	sum := 0
	for _, n := range e.numbers {
		sum += n
	}
	return sum
}

func (e *Equation) computeMultiplication() int {
	sum := 1
	for _, n := range e.numbers {
		sum *= n
	}
	return sum
}
