/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rawdata, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(rawdata), "\n")
	rotation := 50
	rotation_prev := rotation
	zeros := 0
	for _, ins := range data {
		if ins == "" {
			continue
		}
		res, _ := strconv.Atoi(string(ins[1:len(ins)]))
		rotation_prev = rotation
		if string(ins[:1]) == "L" {
			rotation -= res
		} else {
			rotation += res
		}
		for rotation < 0 {
			rotation += 100
			if rotation_prev != 0 && rotation != 0 || rotation < 0 && rotation_prev == 0 {
				zeros += 1
			}
		}
		for rotation > 99 {
			rotation -= 100
			if rotation_prev != 0 && rotation != 0 || rotation < 0 && rotation_prev == 0 {
				zeros += 1
			}
		}
		if rotation == 0 {
			zeros += 1
		}
		// fmt.Println(ins, rotation, zeros)
	}
	fmt.Println("zeros:", zeros)
}
