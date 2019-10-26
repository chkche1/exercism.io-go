package listops

type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

func (it IntList) Append(list IntList) IntList {
	return append(it, list...)
}

func (it IntList) Concat(lists []IntList) IntList {
	res := it
	for _, list := range lists {
		res = append(res, list...)
	}
	return res
}

func (it IntList) Filter(filter predFunc) IntList {
	res := IntList{}
	for _, elem := range it {
		if filter(elem) {
			res = append(res, elem)
		}
	}
	return res
}

func (it IntList) Length() int {
	return len(it)
}

func (it IntList) Map(fn unaryFunc) IntList {
	for idx, elem := range it {
		it[idx] = fn(elem)
	}
	return it
}

func (it IntList) Foldl(fn binFunc, init int) int {
	res := init
	for _, elem := range it {
		res = fn(res, elem)
	}
	return res
}

func (it IntList) Foldr(fn binFunc, init int) int {
	res := init
	for i := len(it) - 1; i >= 0; i-- {
		res = fn(it[i], res)
	}
	return res
}

func (it IntList) Reverse() IntList {
	for i := len(it)/2 - 1; i >= 0; i-- {
		ptr := len(it) - 1 - i
		it[i], it[ptr] = it[ptr], it[i]
	}
	return it
}