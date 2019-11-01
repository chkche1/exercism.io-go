package proverb

import "fmt"

const (
	format = "For want of a %s the %s was lost."
	lastLine = "And all for the want of a %s."
)

/*
Given a list of inputs, generate the relevant proverb.
 */
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return rhyme
	}
	n := len(rhyme)
	res := make([]string, n)
	for i := 1; i < n; i++ {
		res[i-1] = fmt.Sprintf(format, rhyme[i-1], rhyme[i])
	}
	res[n-1] = fmt.Sprintf(lastLine, rhyme[0])
	return res
}
