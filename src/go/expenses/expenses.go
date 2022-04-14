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
func Filter(in []Record, predicate func(Record) bool) (out []Record) {
	for _, rec := range in {
		if predicate(rec) {
			out = append(out, rec)
		}
	}
	return
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		// FWIW, I don't agree with this - seems like it should be inclusive
		// on one end and exclusive on the other - but this is what the tests say!
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p
func TotalByPeriod(in []Record, p DaysPeriod) (total float64) {
	for _, r := range in {
		// There must be `Filter` and `Fold` methods to apply to a slice for this,
		// but I'm intentionally not using them because we haven't been introduced
		// to them yet
		//
		// Update: whoops, we'd implemented `Filter` above. Still surprising that,
		// from a little searching, it doesn't appear to be part of
		// the standard library, though (ref: https://stackoverflow.com/q/37562873/1040915)
		//
		// Keeping my original code for the record, but it would have been more appropriate
		// to do
		//  ```
		//   for _, r := range (Filter(in, ByDaysPeriod(p)) {
		//     total += r.Amount
		//   }
		//   return
		//  ```
		//
		// (And similar below)
		if ByDaysPeriod(p)(r) {
			total += r.Amount
		}
	}
	return
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	// Normally I would just iterate, filter, and accumulate -
	// but we're specifically told to return error if there are no records with that category,
	// and that's not equivalent to "accumulation == 0" (in the presence of
	// non-positive values)
	var filtered_recs []Record
	for _, rec := range in {
		if ByCategory(c)(rec) {
			filtered_recs = append(filtered_recs, rec)
		}
	}
	if len(filtered_recs) == 0 {
		// Note - we hadn't been introduced to error-creation yet?
		return 0, errors.New("filtered records are empty")
	}
	return TotalByPeriod(filtered_recs, p), nil
}
