package stringset

import (
	"fmt"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Set of strings
type Set map[string]bool

// Creates a new empty Set
func New() Set {
	return Set{}
}

// Creates a Set from a slice of strings
func NewFromSlice(l []string) Set {
	set := Set{}
	for _, v := range l {
		set[v] = true
	}
	return set
}

// Print function for the Set
func (s Set) String() string {
	str := ""
	count := 1
	for k := range s {
		str += fmt.Sprintf("%q", k)
		if count < len(s) {
			str += ", "
		}
		count++
	}
	return fmt.Sprintf("{%s}", str)
}

// Returns true is the Set is Empty
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Returns true if set has an element
func (s Set) Has(elem string) bool {
	return s[elem]
}

// Adds an element to the Set
func (s Set) Add(elem string) {
	s[elem] = true
}

// Returns true if all of rhe elements of s1 are contained in s2
func Subset(s1, s2 Set) bool {
	// If the difference is 0 that means that all the elements in s1 are in s2, making s1 a subset of s2
	return len(Difference(s1, s2)) == 0
}

// Returns true if s1 shares no elements with s2
func Disjoint(s1, s2 Set) bool {
	// If intersection is 0 that means the sets have no shared values
	return len(Intersection(s1, s2)) == 0
}

// Returns true if the two sets are equal
func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	set := Union(s1, s2)
	return len(set) == len(s1) && len(set) == len(s2)
}

// Returns a set with the shared values between the two sets
func Intersection(s1, s2 Set) Set {
	set := Set{}
	for k := range s1 {
		if s2[k] {
			set[k] = true
		}
	}
	return set
}

// Returns a set with all the elements that are only in the first set
func Difference(s1, s2 Set) Set {
	set := Set{}
	for k := range s1 {
		if !s2[k] {
			set[k] = true
		}
	}
	return set
}

// Returns a set made of the union of the two sets
func Union(s1, s2 Set) Set {
	set := Set{}
	for k := range s1 {
		set[k] = true
	}
	for k := range s2 {
		set[k] = true
	}
	return set
}

/*
Benchmark using bool for value in map
=== RUN   BenchmarkNewFromSlice1e4
BenchmarkNewFromSlice1e4
BenchmarkNewFromSlice1e4-8          2067            584641 ns/op          373797 B/op        133 allocs/op
PASS
ok      stringset       1.273s

Benchmark using int for value in map
=== RUN   BenchmarkNewFromSlice1e4
BenchmarkNewFromSlice1e4
BenchmarkNewFromSlice1e4-8          1774            631096 ns/op          508281 B/op        243 allocs/op
PASS
ok      stringset       1.194s
*/
