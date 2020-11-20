package utils

import "strings"

// LowerCaseData does a ToLower to the Data string
// of a WorkWith
func LowerCaseData(w WorkWith) WorkWith {
	w.Data = strings.ToLower(w.Data)
	return w
}

// IncrementVersion increments a WorkWith's Version
func IncrementVersion(w WorkWith) WorkWith {
	w.Version++
	return w
}

// OldVersion returns a closures that validates the version
// is greater that the specified amount
func OldVersion(v int) func(w WorkWith) bool {
	return func(w WorkWith) bool {
		return w.Version >= v
	}
}
