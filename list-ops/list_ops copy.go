package listops

import "fmt"

// IntList is an abstraction of a list of integers which we can define methods on
type IntList2 []int

func (s IntList) Foldl2(fn func(int, int) int, initial int) int {
	for _, v := range s {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("%v %T\n", err, err)
				initial = 0
			}
		}()
		initial = fn(v, initial)
	}
	return initial
}

func (s IntList) Foldr2(fn func(int, int) int, initial int) int {
	for _, v := range s.Reverse() {
		defer func() {
			if err := recover(); err != nil {
				initial = 0
			}
		}()
		initial = fn(v, initial)
	}
	return initial
}

func (s IntList) Filter2(fn func(int) bool) IntList {
	filtered := IntList{}
	for _, v := range s {
		if fn(v) {
			filtered = filtered.Append(IntList{v})
		}
	}
	return filtered
}

func (s IntList) Length2() int {
	length := 0
	for range s {
		length++
	}
	return length
}

func (s IntList) Map2(fn func(int) int) IntList {
	mapped := IntList{}
	for _, v := range s {
		mapped = mapped.Append(IntList{fn(v)})
	}
	return mapped
}

func (s IntList) Reverse2() IntList {
	reversed := IntList{}
	for i := len(s) - 1; i >= 0; i-- {
		reversed = reversed.Append(IntList{s[i]})
	}
	return reversed
}

func (s IntList) Append2(lst IntList) IntList {
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

func (s IntList) Concat2(lists []IntList) IntList {
	for _, l := range lists {
		s = s.Append(l)
	}
	return s
}
