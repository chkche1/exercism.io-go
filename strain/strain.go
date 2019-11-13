package strain

type Ints []int
type Lists [][]int
type Strings []string

func (it Ints) Keep(fn func(int) bool) Ints {
	if it == nil {
		return nil
	}
	res := Ints{}
	for _, i := range it {
		if fn(i) {
			res = append(res, i)
		}
	}
	return res
}

func (it Ints) Discard(fn func(int) bool) Ints {
	if it == nil {
		return nil
	}
	res := Ints{}
	for _, i := range it {
		if !fn(i) {
			res = append(res, i)
		}
	}
	return res
}

func (it Lists) Keep(fn func([]int) bool) Lists {
	if it == nil {
		return nil
	}
	res := Lists{}
	for _, i := range it {
		if fn(i) {
			res = append(res, i)
		}
	}
	return res
}

func (it Strings) Keep(fn func(string) bool) Strings {
	if it == nil {
		return nil
	}
	res := Strings{}
	for _, i := range it {
		if fn(i) {
			res = append(res, i)
		}
	}
	return res
}