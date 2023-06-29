package main

import (
	"fmt"
)

func equalizeExpenses(expenses map[string]float64) map[string]float64 {
	totalExpense := 0.0
	for _, expense := range expenses {
		totalExpense += expense
	}

	averageExpense := totalExpense / float64(len(expenses))
	balances := make(map[string]float64)

	for person, expense := range expenses {
		balances[person] = averageExpense - expense
	}

	return balances
}

func main() {
	expenses := map[string]float64{
		"João":   100.0,
		"Maria":  150.0,
		"Carlos": 80.0,
		"André":  120.0,
	}

	balances := equalizeExpenses(expenses)

	fmt.Println("Balances:")
	for person, balance := range balances {
		fmt.Printf("%s: R$%.2f\n", person, balance)
	}
}
