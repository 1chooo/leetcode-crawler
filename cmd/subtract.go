package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Subtract(from string, subtract string) (result string) {
	num1, num2, err := parseTwoFloats(from, subtract)
	if err != nil {
		return ""
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
		diff := Subtract(args[0], args[1])
		if diff == "" {
			fmt.Println("Error: invalid number(s)")
			return
		}
		fmt.Printf("Subtraction of %s from %s = %s.\n\n", args[1], args[0], diff)
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)
}
