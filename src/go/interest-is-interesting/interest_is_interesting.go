package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return 3.213
	case balance >= 0 && balance < 1000:
		return 0.5
	case balance >= 1000 && balance < 5000:
		return 1.621
	default:
		return 2.475
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	// Feels unusual for a Pythonista to be doing this upfront rather than catching
	// an exception - but idiomatic Go appears to eschew Exceptions, so far as I can see
	if (balance <= 0 && targetBalance > 0) ||
		(balance >= 0 && targetBalance < 0) ||
		(balance != 0 && targetBalance == 0) {
		// Undefined what we should return if this never happens
		// (e.g. for a negative value and a positive targetBalance) - returning -1
		// is the usual convention
		return -1
	}

	var years = 0
	for balance < targetBalance {
		years++
		balance = AnnualBalanceUpdate(balance)
	}
	// Candidly - I didn't bother doing math to figure this out _a priori_,
	// I just relied on the test cases to determine whether I was off-by-one.
	// Which is exactly what I'd do in real-life, too - I trust my ability to
	// write explicit tests more than I trust my ability to do edge-case arithmetic
	// (and they have the advantage of being repeatable regression tests, too :P )
	return years

}
