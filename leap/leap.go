// A package that holds a method that determines whether a year is a leap year
package leap

// Given a year, report if it is a leap year.
func IsLeapYear(n int) bool {
	if n%4 != 0 {
		return false
	}

	if n%400 == 0 {
		return true
	} else if n%100 == 0 {
		return false
	}
	return true
}
