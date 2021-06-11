package word

import "unicode"

func IsPalindrome(s string) bool {

	var letters []rune

	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	i := 0
	j := len(letters) - 1

	for i <= j {

		if letters[i] != letters[j] {
			return false
		}

		i++
		j--
	}

	return true
}
