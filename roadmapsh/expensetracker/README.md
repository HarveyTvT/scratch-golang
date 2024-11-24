# How to run

```bash
$ go build expense-tracker.go

$ ./expense-tracker add --description "Lunch" --amount 20
# Expense added successfully (ID: 1)

$ ./expense-tracker add --description "Dinner" --amount 10
# Expense added successfully (ID: 2)

$ ./expense-tracker list
# ID  Date       Description  Amount
# 1   2024-08-06  Lunch        $20
# 2   2024-08-06  Dinner       $10

$ ./expense-tracker summary
# Total expenses: $30

$ ./expense-tracker delete --id 2
# Expense deleted successfully

$ ./expense-tracker summary
# Total expenses: $20

$ ./expense-tracker summary --month 8
# Total expenses for August: $20
```
