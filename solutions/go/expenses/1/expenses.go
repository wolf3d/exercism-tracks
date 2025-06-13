package expenses

import "errors"

// Record represents an expense record.
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
    for i := 0; i < len(in); i++ {
        item := in[i]
		if predicate(item) {
            out = append(out, item)
        }
    }
    return out
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise
func ByDaysPeriod(p DaysPeriod) func(r Record) bool {
	return func(r Record) bool  { 
        return r.Day >= p.From && r.Day <= p.To
    }
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise
func ByCategory(c string) func(r Record) bool {
	return func(r Record) bool  { 
        return r.Category == c
    }
}

func CalcTotal(in []Record) float64 {
    var total float64
    for i := 0; i < len(in); i++ {
        total += in[i].Amount
    }
	return total
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
    var out []Record = Filter(in, ByDaysPeriod(p))
	return CalcTotal(out)
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var out []Record = Filter(in, ByCategory(c))
    if len(out) == 0 {
        return 0.0, errors.New("no records in the list")
    }
	out = Filter(out,ByDaysPeriod(p))
	return CalcTotal(out), nil
}
