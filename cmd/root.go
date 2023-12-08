package cmd

import (
	"fmt"
	"os"

	"github.com/jfpferreira/adventofcode23/cmd/day1"
	"github.com/jfpferreira/adventofcode23/cmd/day2"
	"github.com/spf13/cobra"
)

var AppName string

var rootCmd = &cobra.Command{
	Use: AppName,
	Short: "Solutions for Advent of Code 2023",
}

func Execute()  {
	day1.AddCommands(rootCmd)
	day2.AddCommands(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}