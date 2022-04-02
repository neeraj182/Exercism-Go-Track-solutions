package main

import (
	"fmt"
)

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
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; !ok { // if unit is not in units
		return false
	}
	if _, ok := bill[item]; ok { // if item is already in bill
		bill[item] += units[unit]
	} else {
		bill[item] = units[unit] // add item to bill
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	Item := bill[item]
	if _, ok := bill[item]; !ok {
		return false
	}

	total := units[unit]
	if _, ok := units[unit]; !ok {
		return false
	}
	New := Item - total

	if New < 0 {
		return false
	} else if New == 0 {
		delete(bill, item)
		return true
	} else {
		bill[item] = New
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if _, ok := bill[item]; !ok {
		return 0, false

	}
	total := bill[item]
	return total, true
}

func main() {
	fmt.Println(Units())
	fmt.Println(NewBill())
	bill := map[string]int{"carrot": 12, "grapes": 3}
	units := Units()
	//fmt.Println(AddItem(bill, units, "carrot", "dozen"))
	fmt.Println(RemoveItem(bill, units, "mango", "dozen"))
	fmt.Println(GetItem(bill, "mango"))
}
