package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Add(first string, second string) (result string) {
	num1, num2, err := parseTwoFloats(first, second)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%f", num1+num2)
}

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"addition"},
	Short:   "Add 2 numbers",
	Long:    "Carry out addition operation on 2 numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		sum := Add(args[0], args[1])
		if sum == "" {
			fmt.Println("Error: invalid number(s)")
			return
		}
		fmt.Printf("Addition of %s and %s = %s.\n\n", args[0], args[1], sum)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
