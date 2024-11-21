package ssha

import (
	"math"
)

// This is all very messy, it wasn't meant to be good, much less be readable ;)

func Hash(in string) string {
	var t = [62]rune{ // Table of chars available, for input/output
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
	var out = [32]int{ // Template string to change, 128 bits long
		28, 15, 6, 31, 17, 22, 55, 23, 9,
		10, 13, 25, 3, 26, 19, 11, 4, 29,
		5, 24, 18, 21, 7, 16, 2, 30, 12,
		32, 14, 8, 42, 58,
	}
	inS := make([]int, 0) // Used to simplify table usage
	overlap := len(in) - 32

	// Check if input is in table
	for _, char := range in {
		for j, runeVal := range t {
			if runeVal == char {
				inS = append(inS, j)
				break
			}
		}
	}

	// Hash characters from in
	for i, lenI := 0, len(in); i < lenI; i++ {
		c := float64(inS[i])

		// Math processes
		c *= math.Sin(129) + math.Sqrt(9391)/math.Log10(c)
		c += math.Mod(math.Pow(c, 29)/math.Sqrt(3), 17)
		c = math.Mod(c, 255)

		inS[i] = int(c)
	}

	// Fix overlap before string processing
	if overlap < 0 {
		overlap = 0
	}

	// Combine into out[]
	/* TODO: Loop multiple times to account for length of input
	// Ex: input = jfz
	// l = len(input)
	// loop over loop below (nested?) in order to go over the full
	// table l amount of times
	*/
	for i := 0; i < len(out); i++ {
		for j := 0; j < len(in); j++ {
			if inS[j] < 128 {
				out[i] = (out[i] + inS[j]) % 62
			} else {
				out[i] = (out[i] - inS[j]) % 62
			}
		}
	}

	outS := make([]rune, len(out))
	// Process into rune slice
	for i := 0; i < len(out); i++ {
		outS[i] = t[int(math.Abs(float64(out[i])))]
	}

	output := string(outS)
	return output
}
