package cmd

import (
	"ExpenseTracker/trackerService"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	month      int
	year       int
	day        int
	summaryCmd = &cobra.Command{
		Use:   "summary",
		Short: "summary",
		Long:  `summary of all expenses`,
		Run: func(cmd *cobra.Command, args []string) {
			var expdate int64 = 0
			if month != -1 {
				expdate += int64(month) * 30
			}
			if year != -1 {
				expdate += int64(year) * 12 * 30
			}
			if day != -1 {
				expdate += int64(day)
			}
			expdate *= 24 * 60 * 60 * 1000
			expenses := ExpenseService.GetExpense(func(expense trackerService.Expense) bool {
				ret := true
				if category != "" {
					if expense.Category != category {
						ret = false
					}
				}
				if expdate != 0 {
					if expdate <= expense.Time {
						ret = false
					}
				}
				return ret
			})
			counter := 0
			for _, expense := range expenses {
				counter += expense.Amount
			}
			fmt.Println("# Total expenses: $", counter)
		},
	}
)

func init() {
	RootCmd.AddCommand(summaryCmd)
	summaryCmd.Flags().StringVarP(&category, "category", "c", "", "category")
	summaryCmd.Flags().IntVarP(&month, "month", "m", -1, "month of expenses")
	summaryCmd.Flags().IntVarP(&year, "year", "y", -1, "year of expenses")
	summaryCmd.Flags().IntVarP(&day, "day", "d", -1, "day of expenses")
}
