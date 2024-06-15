package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	filtered := IntList{}
	for _, v := range s {
		if fn(v) {
			filtered = filtered.Append(IntList{v})
		}
	}
	return filtered
}

func (s IntList) Length() int {
	length := 0
	for range s {
		length++
	}
	return length
}

func (s IntList) Map(fn func(int) int) IntList {
	mapped := IntList{}
	for _, v := range s {
		mapped = mapped.Append(IntList{fn(v)})
	}
	return mapped
}

func (s IntList) Reverse() IntList {
	reversed := IntList{}
	for i := len(s) - 1; i >= 0; i-- {
		reversed = reversed.Append(IntList{s[i]})
	}
	return reversed
}

func (s IntList) Append(lst IntList) IntList {
	appended := make([]int, s.Length()+lst.Length())
	index := 0
	for _, v := range s {
		appended[index] = v
		index++
	}
	for _, v := range lst {
		appended[index] = v
		index++
	}
	return appended
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, l := range lists {
		s = s.Append(l)
	}
	return s
}
