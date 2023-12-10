package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type Set map[string]int

type Game struct {
	id int
	sets []*Set
}

func parse(input []string) []*Game {
	games := []*Game{}
	for _, line := range input {
		parts := strings.Split(line, ":")

		i, err := strconv.Atoi(strings.Split(parts[0], " ")[1]) 
		if err != nil {
			panic("The identifier was not a number");
		}
		game := &Game{
			id: i,
			sets: []*Set{},
		}
		
		for _, set := range strings.Split(strings.TrimSpace(parts[1]), ";") {

			currentSet := make(Set)
			for _, cube_set := range strings.Split(strings.TrimSpace(set), ",") {
				fetch := strings.Split(strings.TrimSpace(cube_set), " ")

				number, err := strconv.Atoi(fetch[0])
				if err != nil {
					panic("The number was not")
				}

				currentSet[fetch[1]] = number
			}
			game.sets = append(game.sets, &currentSet)
		}
		games = append(games, game)
	}

	return games
}

var cubes = map[string]int{
	"red": 12,
	"green": 13,
	"blue": 14,
}

func solveA(input []string) string {
	games := parse(input)
	
	total := 0
	for _, game := range games {
		if isGameValid(game) {
			total += game.id
		}	
	}

	return fmt.Sprint(total)
}

func isGameValid(game *Game) bool {
	for _, set := range game.sets {
		for k, v := range *set {
			if v > cubes[k] {
				return false
			}
		}
	}
	return true
}

func solveB(input []string) string {
	games := parse(input)
	total := 0

	for _, game := range games {
		power := 1
		for _, v := range minimumCubes(game) {
			power *= v
		}

		total += power
	}

	return fmt.Sprint(total)
}

func minimumCubes(game *Game) map[string]int {
	minimumCubes := map[string]int{
		"red": 0,
		"green": 0,
		"blue": 0,
	}

	for _, set := range game.sets {
		for k, v := range *set {
			if v > minimumCubes[k] {
				minimumCubes[k] = v
			}
		}
	}
	return minimumCubes
}
