package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	m := make(map[string]int)
	m["quarter_of_a_dozen"] = 3
	m["half_of_a_dozen"] = 6
	m["dozen"] = 12
	m["small_gross"] = 120
	m["gross"] = 144
	m["great_gross"] = 1728

	return m
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	b, exists := units[unit]
	if !exists {
		return false
	}

	if quantity, exists := bill[item]; exists {
		bill[item] = quantity + b
	} else {
		bill[item] = b
	}

	// another answer:
	// bill[item] = bill[item] + b
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	u, ok := units[unit]
	if !ok {
		return false
	}

	if b, ok := bill[item]; ok {
		// assign a temp val to check against
		nq := b - u
		if nq < 0 {
			return false
		}
		if nq == 0 {
			delete(bill, item)
			return true
		}
		bill[item] = nq
		return true
	} else {
		return false
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if b, ok := bill[item]; !ok {
		return 0, false
	} else {
		return b, true
	}
}
