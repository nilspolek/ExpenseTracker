package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Long:  `list all Expenses`,
	Run: func(cmd *cobra.Command, args []string) {
		var sb strings.Builder
		sb.WriteString("# ID\t\t\t\t\tDate\t\t\t\tDescription\tAmount\n")
		for _, expense := range ExpenseService.GetExpenses() {
			sb.WriteString(expense.String() + "\n")
		}
		fmt.Println(sb.String())
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
