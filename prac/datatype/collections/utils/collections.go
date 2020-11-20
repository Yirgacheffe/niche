package utils

// WorkWith is the struct we'll be implementing collections for
type WorkWith struct {
	Data    string
	Version int
}

// Filter is functional filter.
// It takes a list of WorkWith and a WorkWith Function that returns a bool for each "true"
// element we return it to the resultant list
func Filter(ws []WorkWith, f func(w WorkWith) bool) []WorkWith {

	// smallest size
	// of work if len = 0
	result := make([]WorkWith, 0)

	for _, v := range ws {
		if f(v) {
			result = append(result, v)
		}
	}

	return result // re, return new filtered list

}

// Map is a functional map. It takes a list of
// WorkWith and a WorkWith Function that takes a WorkWith and returns a modified
// WorkWith.
// The end result is a list of modified WorkWiths
func Map(ws []WorkWith, f func(w WorkWith) WorkWith) []WorkWith {

	// same length with original
	result := make([]WorkWith, len(ws))

	for i, v := range ws {
		result[i] = f(v)
	}

	return result // re, return new mapped work list

}
