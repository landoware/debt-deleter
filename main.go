package main

type Loan struct {
	name        string
	balance     float64
	rate        float64
	min_payment float64
	due_day     int
}

func newLoan(name string, balance, rate, min_payment float64, due_day int) Loan {
	return Loan{
		name:        name,
		balance:     balance,
		rate:        rate,
		min_payment: min_payment,
		due_day:     due_day,
	}
}

func main() {
	// loanA := newLoan("Nelnet AA", 7500.00, 0.05, 200.0, 18)
	//
	// fmt.Printf("Amortizing %.2f with payments of $%.2f\n\n", loanA.balance, loanA.min_payment)
	//
	// if loanAPayoff, loanASchedule, err := getNumberOfPayments(loanA.balance, loanA.rate, 100, loanA.due_day); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s will pay off in %d payments", loanA.name, loanAPayoff)
	// 	fmt.Printf("\n\n%v", loanASchedule)
	// }
}
