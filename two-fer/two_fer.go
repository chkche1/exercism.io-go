package twofer

import "fmt"

// ShareWith returns "One for you, one for me." when the input string is empty. Otherwise it returns a string referencing the input name
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
