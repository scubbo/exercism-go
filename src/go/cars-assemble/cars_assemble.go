package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// I haven't taken the time to myself to prove to myself whether there exists
	// an x, d, such that floor(floor(x)/d) != floor(x/d). Regardless - this is
	// certainly right, even if there _might_ be an alternative way to do it
	// `int(Calculate(...) / 60`
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	bulkCost := uint(carsCount / 10 * 95_000)
	individualCost := uint(carsCount % 10 * 10_000)
	return bulkCost + individualCost
}
