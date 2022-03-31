package main

import (
	"fmt"
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch { // using switch to determine interest rate
	case balance < 0:
		return 3.213
	case balance < 1000:
		return 0.5
	case balance < 5000:
		return 1.621
	default: // if we use default, return won't cause a error
		return 2.475
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)/100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) (total int) {
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance) 
		total++ // total is a variable, so it will be increased by 1 each time
	}
	return
}

func main() {
	var balance float64 = 200.75
	fmt.Println(InterestRate(balance))
	fmt.Println(Interest(balance))
	fmt.Println(AnnualBalanceUpdate(balance))
	fmt.Println(YearsBeforeDesiredBalance(balance, 214.88))
}
