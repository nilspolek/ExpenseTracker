package cmd

import (
	"ExpenseTracker/trackerService"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	ExpenseService trackerService.ExpensesService
)

var RootCmd = &cobra.Command{
	Use:   "Expense Tracker",
	Short: "Expense Tracker",
	Long:  "A simple Expense Tracker",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Expense Tracker Use --help to see available commands.")
	},
}

func Execute() {
	ExpenseService, err = trackerService.NewExpensesService()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
