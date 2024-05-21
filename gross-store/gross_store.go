package gross

var units map[string]int = map[string]int{
	"quarter_of_a_dozen": 3,
	"half_of_a_dozen":    6,
	"dozen":              12,
	"small_gross":        120,
	"gross":              144,
	"great_gross":        1728,
}

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// ItemExists checks it item exists in map
func ItemExists(key string, map_ map[string]int) bool {
	_, ok := map_[key]
	return ok
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if !ItemExists(unit, units) {
		return false
	}
	bill[item] += units[unit]
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if !ItemExists(item, bill) || !ItemExists(unit, units) {
		return false
	}
	new_quantity := bill[item] - units[unit]
	if new_quantity < 0 {
		return false
	}
	if new_quantity == 0 {
		delete(bill, item)
	} else {
		bill[item] = new_quantity
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	v, ok := bill[item]
	return v, ok
}
