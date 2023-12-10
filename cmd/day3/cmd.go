package day3

import (
	"github.com/jfpferreira/adventofcode23/cmd/challenge"
	"github.com/spf13/cobra"
)

func AddCommands(cmd *cobra.Command) {
	daySolution := &challenge.DaySolution{
		Day: 3,
		SolveA: solveA,
		SolveB: solveB,
	}

	daySolution.AddDaySolutions(cmd)
}
