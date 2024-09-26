package cmd

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	uid       uuid.UUID
	id        string
	ammount   int
	err       error
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "update",
		Long:  `update an Expense`,
		Run: func(cmd *cobra.Command, args []string) {
			uid, err = uuid.Parse(id)
			if err != nil {
				fmt.Println("error: ", err)
			}
			err := ExpenseService.UpdateExpense(uid, ammount)
			if err != nil {
				fmt.Println("error: ", err)
			}
		},
	}
)

func init() {
	RootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&id, "id", "i", "", "Expense ID")
	updateCmd.Flags().IntVarP(&ammount, "ammount", "a", 1, "Expense Amount")
}
