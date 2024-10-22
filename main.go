package main

import (
	"fmt"
	"math"
	"strconv"
)

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

var out = [32]uint8{
	28, 15, 6, 31, 17, 22, 55, 23, 9,
	10, 13, 25, 3, 26, 19, 11, 4, 29,
	5, 24, 18, 21, 7, 16, 2, 30, 12,
	32, 14, 8, 42, 58,
}

func main() {
	var current float64 = 24
	var nextChar float64 = 38

	// Math shit
	current *= math.Sin(129) + math.Sqrt(9391)/math.Log10(nextChar)
	current += math.Mod(math.Pow(nextChar, 29)+math.Sqrt(current/4), 2901)
	current = math.Mod(current, 62)
	current = math.Abs(current)

	var current2 int
	current2 = int(math.Mod(current, 62))

	fmt.Println(strconv.Itoa(current2))
}
