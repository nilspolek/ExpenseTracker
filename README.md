# Expense Tracker - README

## Introduction
The Expense Tracker is a command-line application designed to help users track their expenses and manage their finances. This simple yet effective tool allows users to add, update, delete, and view expenses, as well as generate summaries for specific time periods. This project is perfect for those looking to build a command-line application that manipulates and manages data.

## Features
The application provides the following core functionalities:

1. **Add Expense**: Users can add an expense by providing a description and amount.
2. **Update Expense**: Users can update existing expenses based on the expense ID.
3. **Delete Expense**: Users can delete an expense using its ID.
4. **View All Expenses**: Users can view a list of all recorded expenses.
5. **View Expense Summary**: Users can see the total expenses recorded.
6. **View Monthly Summary**: Users can view expenses specific to a particular month of the current year.

### Additional Features
The application can be extended with the following functionalities:

1. **Expense Categories**: Users can categorize expenses and filter them by category.
2. **Budget Management**: Users can set a budget for each month, with alerts if the budget is exceeded.
3. **Export to CSV**: Users can export the list of expenses to a CSV file for easy sharing and further analysis.

## Usage
The application should be run from the command line with various commands and options. Below is a summary of available commands and their usage:

### Commands
1. **Add Expense**
    ```bash
    $ expense-tracker add --description "Lunch" --amount 20
    ```
    Output:  
    `Expense added successfully (ID: 1)`

2. **Update Expense**
    ```bash
    $ expense-tracker update --id 1 --amount 25
    ```
    Output:  
    `Expense updated successfully (ID: 1)`

3. **List Expenses**
    ```bash
    $ expense-tracker list
    ```
    Output:  
    ```
    ID   Date          Description  Amount
    1    2024-08-06    Lunch        $20
    2    2024-08-06    Dinner       $10
    ```

4. **Delete Expense**
    ```bash
    $ expense-tracker delete --id 1
    ```
    Output:  
    `Expense deleted successfully`

5. **View Total Summary**
    ```bash
    $ expense-tracker summary
    ```
    Output:  
    `Total expenses: $30`

6. **View Monthly Summary**
    ```bash
    $ expense-tracker summary --month 8
    ```
    Output:  
    `Total expenses $20`

## Contributing
If you would like to contribute to this project, feel free to fork the repository and submit a pull request with your changes. Make sure to include a detailed description of what you have modified.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements
Thanks to everyone who provided feedback and suggestions for improving this project. This project idea is a great starting point for anyone looking to build their first command-line application.

---

Enjoy tracking your expenses and stay financially healthy! 💸
