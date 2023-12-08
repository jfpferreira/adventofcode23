package day1

import (
	"strconv"
	"testing"
)

func TestA(t *testing.T){
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	expected := 142

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
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	expected := 281

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