package main

import "fmt"

// lc 151 : reverse the order of the words in a given string, remove space at the beginning / end and double space
// space O(n), in c++ : use dp + resize
func WordRev(input []string) []string {

	/*first reverse : the entire string*/
	for i := 0; i < len(input); i++ {
		if i > len(input) - 1 - i {
			break
		}
		buff := input[i]
		input[i] = input[len(input) - 1 - i]
		input[len(input) - 1 - i] = buff
	}

	/*remove the space*/
	fP, sP := 0, 0
	for len(input) > 0 && input[fP] == " " && fP < len(input) { // remove the redundant space at the start
		fP++
	}

	for ;fP < len(input); fP++ {
		// remove the redundant space in the middle : continue
		if fP > 1 && input[fP] == input[fP - 1] && input[fP] == " " {
			continue
		}

		input[sP] = input[fP]
		sP++
	}

fmt.Println(input, sP)

	// reverse every word
	lengthReverse := 0
	for i := 0; i < sP; i++ {
		if input[i] == " " || i == sP - 1 {
			count := 0
			for j := i - lengthReverse; j < i; j++ {
				count ++
				if j > i - count {
					break
				}
				buff := input[j]
				input[j] = input[i - count]
				input[i - count] = buff
			}
			lengthReverse = 0
		} else {
			lengthReverse++

		}
	}


	// remove the redundant space at the end : control the length of HEAD->sP and the caracter of sP
	if sP == len(input) {		// no changes
		return input
	}
	if sP - 1 > 0 && input[sP - 1] == " " {
		return input[:sP - 1]

	} else {
		return input[:sP]
	}


}