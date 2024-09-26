package trackerService

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"os"
	"time"
)

// Expense represents a single expense record.
type Expense struct {
	Id          uuid.UUID // Unique identifier for the expense
	Description string    // Description of the expense
	Amount      int       // Amount spent
	Time        int64     // Timestamp of the expense
	Category    string
}

// ExpenseFromConfig is used for reading expenses from the configuration.
type ExpenseFromConfig struct {
	Id          string // UUID as a string
	Description string
	Amount      int
	Time        int64
	Category    string
}

// ExpensesService manages a collection of expenses.
type ExpensesService struct {
	expenses []Expense // Slice to hold multiple Expense records
}

// Error messages.
var (
	ErrExpenseNotFound = errors.New("expense not found")
)

// NewExpensesService initializes and returns a new ExpensesService with data loaded from config.
func NewExpensesService() (ExpensesService, error) {
	// Ensure config file exists
	if err := ensureConfigFile(".expenses.json"); err != nil {
		return ExpensesService{}, err
	}

	// Read configuration
	if err := viper.ReadInConfig(); err != nil {
		return ExpensesService{}, err
	}

	// Unmarshal expenses from configuration
	var expensesFromCfg []ExpenseFromConfig
	if err := viper.UnmarshalKey("expenses", &expensesFromCfg); err != nil {
		return ExpensesService{}, err
	}

	// Convert to internal Expense type
	expenses := convertToExpenses(expensesFromCfg)
	return ExpensesService{expenses: expenses}, nil
}

// ensureConfigFile checks if the config file exists; if not, creates a new one.
func ensureConfigFile(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("Config file does not exist, creating a new one...")
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString("{}"); err != nil {
			return err
		}
	}
	return nil
}

// convertToExpenses transforms a slice of ExpenseFromConfig into a slice of Expense.
func convertToExpenses(expenseCfg []ExpenseFromConfig) []Expense {
	expenses := make([]Expense, len(expenseCfg))
	for i, cfg := range expenseCfg {
		id, err := uuid.Parse(cfg.Id)
		if err != nil {
			continue // Skip invalid UUIDs
		}
		expenses[i] = Expense{
			Id:          id,
			Description: cfg.Description,
			Amount:      cfg.Amount,
			Time:        cfg.Time,
			Category:    cfg.Category,
		}
	}
	return expenses
}

// AddExpense adds a new expense to the service and updates the config.
func (es *ExpensesService) AddExpense(expense Expense) error {
	expense.Id = uuid.New()
	expense.Time = time.Now().Unix()
	es.expenses = append(es.expenses, expense)

	// Update configuration with the new expense
	viper.Set("expenses", es.expenses)
	if err := viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}

// GetExpenses returns a list of all stored expenses.
func (es *ExpensesService) GetExpenses() []Expense {
	return es.expenses
}

// GetTotalExpenses calculates and returns the total amount spent.
func (es *ExpensesService) GetTotalExpenses() int {
	total := 0
	for _, expense := range es.expenses {
		total += expense.Amount
	}
	return total
}

// GetExpense filters and returns expenses matching the given criteria.
func (es *ExpensesService) GetExpense(filter func(expense Expense) bool) []Expense {
	var result []Expense
	for _, expense := range es.expenses {
		if filter(expense) {
			result = append(result, expense)
		}
	}
	return result
}

// DeleteExpense removes an expense by its ID.
func (es *ExpensesService) DeleteExpense(id uuid.UUID) error {
	for i, expense := range es.expenses {
		if expense.Id == id {
			es.expenses = append(es.expenses[:i], es.expenses[i+1:]...)
			return es.saveExpenses()
		}
	}
	return ErrExpenseNotFound
}

// UpdateExpense updates the amount of an expense identified by its ID.
func (es *ExpensesService) UpdateExpense(id uuid.UUID, amount int) error {
	for i, expense := range es.expenses {
		if expense.Id == id {
			es.expenses[i].Amount = amount
			return es.saveExpenses()
		}
	}
	return ErrExpenseNotFound
}

// saveExpenses writes the current state of expenses to the configuration file.
func (es *ExpensesService) saveExpenses() error {
	viper.Set("expenses", es.expenses)
	return viper.WriteConfig()
}

// String returns a formatted string representation of the expense.
func (expense *Expense) String() string {
	return fmt.Sprintf("# %s\t%s\t%s\t\t$%d", expense.Id, time.Unix(expense.Time, 0), expense.Description, expense.Amount)
}
