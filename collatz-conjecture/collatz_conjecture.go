// Package that is home for methods related to Collatz Conjecture
package collatzconjecture

import "errors"

/*
Given a number n, return the number of steps required to reach 1.
 */
func CollatzConjecture(n int) (int, error)  {
	if n < 1 {
		return -1, errors.New("input is less than 1")
	}
	res := 0
	for ; n != 1; {
		if n % 2 == 0 {
			n /= 2
		} else {
			n = n * 3 + 1
		}
		res++
	}
	return res, nil
}