package main

import "fmt"

func main() {
	str := "ABCDE"
	length := len(str)
	track := make([]bool, length)

	permutation(str, "", track)
}

// In order go generate permutation, two things are needed: what characters are
// used in this permutation, and what characters have been used previously.

// These 2 variables looks like same, but it's different. For example,
// already generated: ABCD, for AB__ is fixed, C cannot be third, D cannot be forth
// in process: AB__, C & D can be used

// This implementation is really beautiful, is uses sequentially increment index
// to denote any number is already used historically (not in this round), and
// the other array to show some characters is used in the round.

// Probably because I am not familiar with the process, it takes me 3-4 days to
// understand the meaning by hand calculation.

func permutation(str, result string, track []bool) {
	for i := range track {
		if track[i] {
			continue
		}

		track[i] = true
		permutation(str, result+string(str[i]), track)
		track[i] = false
	}

	if len(result) == len(str) {
		fmt.Println(result)
	}
}
