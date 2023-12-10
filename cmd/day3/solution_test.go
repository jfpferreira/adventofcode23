package day3

import (
	"strconv"
	"testing"
)

func TestA(t *testing.T){
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	expected := 4361

	result := solveA(input)
	i, err := strconv.Atoi(result) 
	if err != nil {
		t.Fatal("Result was not an int");
		return
	}

	if i != expected{
		t.Fatalf("Expected: %d, Got: %d", expected, i)
		return
	}
}

func TestB(t *testing.T){
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	expected := 467835

	result := solveB(input)
	i, err := strconv.Atoi(result) 
	if err != nil {
		t.Fatal("Result was not an int");
		return
	}

	if i != expected{
		t.Fatalf("Expected: %d, Got: %d", expected, i)
		return
	}
}