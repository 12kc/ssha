package main

import (
	"fmt"
	"math"
)

var (
	t = [62]rune{
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
	out = [32]uint8{
		28, 15, 6, 31, 17, 22, 55, 23, 9,
		10, 13, 25, 3, 26, 19, 11, 4, 29,
		5, 24, 18, 21, 7, 16, 2, 30, 12,
		32, 14, 8, 42, 58,
	}
	input   string
	overlap int = len(input) - 32
)

func main() {
	in := make([]int, 0)

	fmt.Println("Input the string to hash.")
	fmt.Scan(&input)

	// Check if input is in table
	for _, char := range input {
		for j, runeVal := range t {
			if runeVal == char {
				in = append(in, j)
				break
			}
		}
	}

	for i, l := 0, len(in); i < l+overlap; i++ {
		c := hash(float64(in[i]), float64(in[i+1]))
		out[i] = (out[i] + uint8(c)) % 62
	}
	fmt.Println(out)
}

func hash(current float64, nextChar float64) int {
	current *= math.Sin(129) + math.Sqrt(9391)/math.Log10(nextChar)
	current += math.Mod(math.Pow(nextChar, 29)+math.Sqrt(current/4), 2901)
	current = math.Mod(current, 62)
	return int(current)
}
