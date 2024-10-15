package main

// May God have mercy on those who decide to read this
import (
	"fmt"
	"math"
)

// Table of chars used in final hash
var t = [62]rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
	'i', 'j', 'k', 'l', 'm', 'n', 'o',
	'p', 'q', 'r', 's', 't', 'u', 'v',
	'w', 'x', 'y', 'z', 'A', 'B', 'C',
	'D', 'E', 'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O', 'P', 'Q',
	'R', 'S', 'T', 'U', 'V', 'W', 'X',
	'Y', 'Z', '1', '2', '3', '4', '5',
	'6', '7', '8', '9', '0',
}
var input string
var output []byte

func main() {
	fmt.Scan(&input)
	output = []byte(input)

	for len, c := 0, 1; c <= len; c++ {
		var current rune
		var currentNum float64
		current = rune(input[c])

		// Check location of char in array
		for i, aLen := 0, cap(t); i < aLen; i++ {
			if t[i] == current {
				currentNum = float64(i)
			}
		}

		// Math shit
		currentNum *= math.SqrtPi + 238295/math.Sqrt(7)
		currentNum /= 2949 ^ 3
		currentNum *= math.Acos(currentNum)
		currentNum = math.Mod(currentNum, 62)

		current = t[int(currentNum)]
		output[c] = byte(current)
		fmt.Println(output)
	}
}
