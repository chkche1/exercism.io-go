/*
Package that holds functions for the Darts game
 */
package darts

/*
Write a function that returns the earned points in a single toss of a Darts game
 */
func Score(x, y float64) int {
	r := x * x + y * y
	switch {
	case r <= 1:
		return 10
	case r <= 25:
		return 5
	case r <= 100:
		return 1
	default:
		return 0
	}
}
