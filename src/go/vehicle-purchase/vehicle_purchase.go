package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	// In real life we'd do (something like
	// `return []string{"car", "truck"}.contains(kind)`
	// but I'm still learning, and don't know how to do
	// that idiomatically.
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	// No ternary operator? :'(
	var choice string
	if option1 <= option2 {
		choice = option1
	} else {
		choice = option2
	}
	return choice + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * 0.8
	} else if age < 10 {
		return originalPrice * 0.7
	} else {
		return originalPrice * 0.5
	}
}
