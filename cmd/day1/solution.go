package day1

import (
	"fmt"
	"strings"
)

func solveA(input []string) string {
	total := 0
	for _, line := range input {
		runes := []rune(line)
		for i := 0; i < len(runes); i++ {
			if v, ok := getDigit(runes[i]) ; ok {
				total += v * 10
				break
			}
		}

		for i := len(runes)-1; i >= 0; i-- {
			if v, ok := getDigit(runes[i]) ; ok {
				total += v
				break
			}
		}
	}

	return fmt.Sprint(total)
}

var numbers = map[string]int {
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
}

func solveB(input []string) string {
	total := 0
	for _, line := range input {
		first, last := len(line), -1
		firstV, lastV := 0, 0 
		for k, v := range numbers {
			if l := strings.LastIndex(line, k) ; l != -1 {
				if last < l {
					last = l
					lastV = v
				}
			}
			
			if f := strings.Index(line, k) ; f != -1 {
				if first > f {
					first = f
					firstV = v
				}
			}
		}

		runes := []rune(line)
		for i := 0; i < first; i++ {
			if v, ok := getDigit(runes[i]) ; ok {
				firstV = v
				break
			}
		}

		for i := len(runes)-1; i > last; i-- {
			if v, ok := getDigit(runes[i]) ; ok {
				lastV = v
				break
			}
		}

		total += (firstV*10) + lastV
	}
	return fmt.Sprint(total)
}

func getDigit(i rune) (int, bool) {
	if '0' <= i && i <= '9' {
		return int(i - '0'), true
	}
	return -1, false
}