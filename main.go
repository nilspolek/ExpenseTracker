package main

import (
	"ExpenseTracker/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName(".expenses.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	cmd.Execute()
}
