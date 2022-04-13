package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	// Or, `return map[string]int{}`
	// Or, change signature to `func NewBill() (bill map[string]int)` and body is just `return`)
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, exists := units[unit]
	if !exists {
		return false
	}
	bill[item] += unitValue
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitValue, unitExists := units[unit]
	if !unitExists {
		return false
	}

	curVal, itemExists := bill[item]
	if !itemExists {
		return false
	}

	decrementedValue := curVal - unitValue
	switch {
	case decrementedValue < 0:
		return false
	case decrementedValue == 0:
		delete(bill, item)
		return true
	default:
		bill[item] = decrementedValue
		return true
	}

	// Could instead have skipped the `return true` in the final
	// two cases above and done a `return true` here as a fall-through -
	// but the approach above feels more explicit
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	val, exists := bill[item]
	if !exists {
		return 0, false
	}
	return val, true
}
