package day3

import (
	"fmt"
	"strconv"
)

type Blueprint struct {
	layout []string
	height int
	width int
}

func solveA(input []string) string {
	bp := &Blueprint{
		layout: input,
		height: len(input),
		width: len(input[0]),
	}
	
	total := 0
	for index, line := range input {
		chars := []rune(line)
		prev, next := 0, 0
		for next < bp.width {
			if !isDigit(chars[next]) {
				if prev != next {
					if number, ok := isValidNumber(bp, prev, next-1, index) ; ok{
						total += number
					}
					prev = next
				}
				prev++
			}
			next++
		}

		if prev != next {
			if number, ok := isValidNumber(bp, prev, next-1, index) ; ok{
				total += number
			}
		}
	}
	return fmt.Sprint(total)
}

func isValidNumber(bp *Blueprint, start int, end int, row int) (int, bool) {
	number, err := strconv.Atoi(bp.layout[row][start:end+1])
	if err != nil {
		panic("It is not a number")
	}

	if isSymbol(bp, start-1, row) || isSymbol(bp, end+1, row) {
		return number, true
	}

	for i := start-1; i <= end +1; i++ {
		if isSymbol(bp, i, row-1) || isSymbol(bp, i, row+1) {
			return number, true
		}
	}

	return -1, false
}

func isSymbol(bp *Blueprint, x int, y int) bool {
	if x < 0 || x >= bp.width || y < 0 || y >= bp.height {
		return false
	}

	char := []rune(bp.layout[y])[x]
	return char != '.' 
}

type Position struct {
	x int
	y int
}
func solveB(input []string) string {
	bp := &Blueprint{
		layout: input,
		height: len(input),
		width: len(input[0]),
	}
	
	asterisks := make(map[Position][]int)
	for index, line := range input {
		chars := []rune(line)
		prev, next := 0, 0
		for next < bp.width {
			if !isDigit(chars[next]) {
				if prev != next {
					if number, found, ok := isAdjacent(bp, prev, next-1, index) ; ok{
						for _, pos := range found {
							if _, ok := asterisks[pos] ; !ok {
								asterisks[pos] = []int{}
							}
							asterisks[pos] = append(asterisks[pos], number)
						}
					}
					prev = next
				}
				prev++
			}
			next++
		}

		if prev != next {
			if number, found, ok := isAdjacent(bp, prev, next-1, index) ; ok{
				for _, pos := range found {
					if _, ok := asterisks[pos] ; !ok {
						asterisks[pos] = []int{}
					}
					asterisks[pos] = append(asterisks[pos], number)
				}
			}
		}
	}

	total := 0
	for _, asterisk := range asterisks {
		if len(asterisk) == 2 {
			total += asterisk[0] * asterisk[1]
		}
	}

	return fmt.Sprint(total)
}

func isAdjacent(bp *Blueprint, start int, end int, row int) (int, []Position, bool) {
	number, err := strconv.Atoi(bp.layout[row][start:end+1])
	if err != nil {
		panic("It is not a number")
	}

	asterisks := []Position{}

	if isAsterisk(bp, start-1, row) {
		asterisks = append(asterisks, Position{
			x: start-1,
			y: row,
		})
	}

	if isAsterisk(bp, end+1, row) {
		asterisks = append(asterisks, Position{
			x: end+1,
			y: row,
		})
	}

	for i := start-1; i <= end +1; i++ {
		if isAsterisk(bp, i, row-1) {
			asterisks = append(asterisks, Position{
				x: i,
				y: row-1,
			})
		}

		if isAsterisk(bp, i, row+1) {
			asterisks = append(asterisks, Position{
				x: i,
				y: row+1,
			})
		}
	}

	return number, asterisks, len(asterisks) > 0
}

func isAsterisk(bp *Blueprint, x int, y int) bool {
	if x < 0 || x >= bp.width || y < 0 || y >= bp.height {
		return false
	}

	return []rune(bp.layout[y])[x] == '*' 
}

func isDigit(i rune) bool {
	return '0' <= i && i <= '9'
}
