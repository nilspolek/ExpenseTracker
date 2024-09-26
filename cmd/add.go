package cmd

import (
	"ExpenseTracker/trackerService"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	description string
	amount      int
	category    string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		if description == "" {
			fmt.Println("No description provided")
			os.Exit(1)
		}
		if amount == 0 {
			fmt.Println("No amount provided")
			os.Exit(1)
		}
		err := ExpenseService.AddExpense(trackerService.Expense{Amount: amount, Description: description, Category: category})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Give a description")
	addCmd.Flags().StringVarP(&description, "category", "c", "", "Give a category")
	addCmd.Flags().IntVarP(&amount, "amount", "a", 0, "Give a amount")
}
