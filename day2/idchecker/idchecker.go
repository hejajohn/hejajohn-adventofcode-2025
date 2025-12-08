package idchecker

import (
	"fmt"
	"math"
	"strconv"
)

func CheckIdRange(startId string, endId string) (int, int) {
	invalidIds := 0
	invalidIdsSum := 0

	startNum := stringToNum(startId)
	endNum := stringToNum(endId)
	numSize := findNumSize(startNum)
	var numSizePow int = int(math.Pow10(numSize))
	var numSizePowMid int = int(math.Pow10(numSize / 2))

	for num := startNum; num <= endNum; num++ {
		if num >= numSizePow {
			numSize += 1
			numSizePow = int(math.Pow10(numSize))
			numSizePowMid = int(math.Pow10(numSize / 2))
		}
		if numSize%2 == 1 {
			continue
		}
		if !checkId(num, numSizePowMid) {
			fmt.Printf("Id %d is invalid\n", num)
			invalidIds += 1
			invalidIdsSum += num
		}
	}

	return invalidIds, invalidIdsSum
}

func checkId(num int, numSizePowMid int) bool {
	lefthalf := num / numSizePowMid
	righthalf := num - (lefthalf * numSizePowMid)
	return lefthalf != righthalf
}

func stringToNum(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}

func findNumSize(number int) int {
	result := 0
	for number > 0 {
		result += 1
		number /= 10
	}
	return result
}
