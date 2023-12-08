package challenge

import (
	"bufio"
	"os"
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/spf13/cobra"
)

type DaySolution struct {
	Day int
	SolveA func([]string) string
	SolveB func([]string) string
}

func (solution *DaySolution) AddDaySolutions(rootCmd *cobra.Command) {
	dayCmd := &cobra.Command{
		Use: fmt.Sprintf("%d", solution.Day),
		Short: fmt.Sprintf("Solutions for day %d", solution.Day),
	}
	
	dayCmd.AddCommand(&cobra.Command{
		Use: "A",
		Short: fmt.Sprintf("Solution for day %d, problem A", solution.Day),
		Run: runDay(solution.Day, solution.SolveA),
	})

	dayCmd.AddCommand(&cobra.Command{
		Use: "B",
		Short: fmt.Sprintf("Solution for day %d, problem B", solution.Day),
		Run: runDay(solution.Day, solution.SolveB),
	})

	rootCmd.AddCommand(dayCmd)
}

func runDay(day int, solve func([]string) string) func(*cobra.Command, []string) {
	return func (_ *cobra.Command, _ []string) {
		input := loadInput(day)
		start := time.Now()
	
		result := solve(input)

		fmt.Printf("Total elapsed time(ms) is: %d\n", time.Since(start).Milliseconds())
		fmt.Printf("Result: %s\n", result)
	}
}

func loadInput(day int) []string{
	 _, callingFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("There was a problem loading the input")
	}

	path := filepath.Dir(callingFile)
	path = filepath.Dir(path)
	path = filepath.Join(path, fmt.Sprintf("day%d", day))
	path = filepath.Join(path, "input.txt")
	fmt.Println(path)

	inputFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	
	reader := bufio.NewScanner(inputFile)
	reader.Split(bufio.ScanLines)

	input := []string{}
	for reader.Scan() {
		input = append(input, reader.Text())
	}

	inputFile.Close()
	return input
}