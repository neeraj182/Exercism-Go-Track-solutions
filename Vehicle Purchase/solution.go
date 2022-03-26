package purchase
// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" {
		return true
	} else if kind == "truck"{
		return true
	}
	return false
}
// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	if option1 > option2 {
		s := option2 + " is clearly the better choice."
		return s
	}
	n := option1 + " is clearly the better choice."
	return n
}
// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		n := originalPrice * 80 / 100
		return n
	} else if age >= 10 {
		s := originalPrice * 50 / 100
		return s
	} else if age >= 3 && age < 10 {
		m := originalPrice * 70 / 100
		return m
	}
	return 0
}
