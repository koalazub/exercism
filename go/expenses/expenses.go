package expenses

import "fmt"

type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var out []Record
	for _, i := range in {
		if predicate(i) {
			out = append(out, i)
		}
	}

	if out != nil {
		return out
	}

	return nil
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {

	return func(r Record) bool {
		if r.Day >= p.From && r.Day <= p.To {
			return true
		}
		return false
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		if c == r.Category {
			return true
		}
		return false
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	total := 0.0
	filter := Filter(in, ByDaysPeriod(p))
	for _, filter := range filter {
		if ByDaysPeriod(p)(filter) {
			total += filter.Amount
		}
	}
	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
// Implement the `CategoryExpenses` function that returns the total amount of expenses in a category in a given period of days. The function should also differentiate the case when the given category is not present in the expenses records and the case when there are no category's expenses in the provided period.
// When the category is not a category of any of the records (regardless of period of time) the function should return an error.

func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var exp float64

	fltRec := Filter(in, ByCategory(c))
	if fltRec == nil {
		return 0, fmt.Errorf("unknown category %s", c)
	}

	exp = TotalByPeriod(fltRec, p)

	return exp, nil
}
