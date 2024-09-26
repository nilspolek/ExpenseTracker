package cmd

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"os"
)

var (
	idToDelete string
	deleteCmd  = &cobra.Command{
		Use:   "delete",
		Short: "delete",
		Long:  `delete a Expense`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(idToDelete)
			if idToDelete == "" {
				fmt.Println("Error: You must provide a Expense ID")
				os.Exit(1)
			}
			uid, err := uuid.Parse(idToDelete)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			err = ExpenseService.DeleteExpense(uid)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			fmt.Println("# Expense deleted successfully")
		},
	}
)

func init() {
	RootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&idToDelete, "id", "i", "", "Expense ID")
}
