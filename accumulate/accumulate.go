package accumulate

/*
Given a collection and an operation to perform on each element of the collection, returns a new collection containing the result of applying that operation to each element of the input collection.
 */
func Accumulate(input []string, fn func(elem string) string) []string  {
	res := make([]string, len(input))
	for i, elem := range input {
		res[i] = fn(elem)
	}
	return res
}