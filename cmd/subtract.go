package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

func Subtract(from string, subtract string) (result string) {
	num1, err := strconv.ParseFloat(from, 64)
	if err != nil {
		fmt.Println("Error: First value is invalid")
		return
	}
	num2, err := strconv.ParseFloat(subtract, 64)
	if err != nil {
		fmt.Println("Error: Second value is invalid")
		return
	}
	return fmt.Sprintf("%f", num1-num2)
}

var subtractCmd = &cobra.Command{
	Use:     "subtract",
	Aliases: []string{"sub"},
	Short:   "Subtract a number from another",
	Long:    "Carry out subtraction operation on 2 integers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Subtraction of %s from %s = %s.\n\n", args[1], args[0], Subtract(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)
}
