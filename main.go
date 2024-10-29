package main

import (
	"fmt"

	ss "github.com/12kc/ssha/pkg/ssha"
)

func main() {
	var input string
	fmt.Println("Input the string to hash.")
	fmt.Scan(&input)
	out := ss.Hash(input)
	fmt.Println(out)
}
