package day4

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Card struct{
	winningNumber []int
	numbers []int
	total int
}

func parse(input []string) []*Card {
	cards := []*Card{}
	for _, cardLine := range input {
		card := &Card{
			winningNumber: []int{},
			numbers: []int{},
			total: 1,
		}
		unreadyCards := strings.Split(strings.TrimSpace(strings.Split(cardLine, ":")[1]), "|")
		for _, number := range strings.Split(strings.TrimSpace(unreadyCards[0]), " ") {
			n, err := strconv.Atoi(number)
			if err != nil {
				continue
			}
			card.winningNumber = append(card.winningNumber, n)
		}

		for _, number := range strings.Split(strings.TrimSpace(unreadyCards[1]), " ") {
			n, err := strconv.Atoi(number)
			if err != nil {
				continue
			}
			card.numbers = append(card.numbers, n)
		}
		sort.Ints(card.winningNumber)
		sort.Ints(card.numbers)
		cards = append(cards, card)
	} 

	return cards
}

func solveA(input []string) string {
	cards := parse(input)
	total := 0.0
	for _, card := range cards {
		local := -1
		wn, n := 0, 0
		for wn < len(card.winningNumber) && n < len(card.numbers) {
			if card.numbers[n] == card.winningNumber[wn] {
				local++
				wn++
				n++
			} else if card.numbers[n] > card.winningNumber[wn] {
				wn++
			} else {
				n++
			}
		}

		if local != -1 {
			total += math.Pow(2, float64(local))
		}
	}

	return fmt.Sprint(int(total))
}

func solveB(input []string) string {
	cards := parse(input)
	total := 0
	for i, card := range cards {
		matches := 0
		wn, n := 0, 0
		for wn < len(card.winningNumber) && n < len(card.numbers) {
			if card.numbers[n] == card.winningNumber[wn] {
				if matches + 1 + i >= len(cards) {
					break
				}
				wn++
				n++
				matches++
				cards[matches + i].total += card.total
			} else if card.numbers[n] > card.winningNumber[wn] {
				wn++
			} else {
				n++
			}
		}
		total += card.total
	}
	return fmt.Sprint(total)
}
