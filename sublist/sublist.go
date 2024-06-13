package sublist

import (
	"strconv"
	"strings"
)

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {

	if len(l1) == 0 && len(l2) > 0 {
		return RelationSublist
	}

	if len(l2) == 0 && len(l1) > 0 {
		return RelationSuperlist
	}

	diff := len(l1) - len(l2)

	if diff == 0 {
		for i, v := range l1 {
			if l2[i] != v {
				return RelationUnequal
			}
		}
		return RelationEqual
	}

	bigger := l1
	smaller := l2

	if diff < 0 {
		bigger = l2
		smaller = l1
	}

	var subset bool

	for i := 0; i+len(smaller) <= len(bigger); i++ {
		subset = true
		for j, v := range bigger[i : i+len(smaller)] {
			if v != smaller[j] {
				subset = false
				break
			}
		}

		if subset {
			if diff < 0 {
				return RelationSublist
			}
			return RelationSuperlist
		}
	}

	return RelationUnequal

}

/*
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^BenchmarkSublist$ sublist
Benchmark
=== RUN   BenchmarkSublist
BenchmarkSublist
BenchmarkSublist-8       8207434               146.3 ns/op             0 B/op          0 allocs/op
PASS
ok      sublist 1.353s


Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^BenchmarkSublist2$ sublist
Benchmark converting to string and checking relation
=== RUN   BenchmarkSublist2
BenchmarkSublist2
BenchmarkSublist2-8       247075              4593 ns/op             720 B/op        111 allocs/op
PASS
ok      sublist 1.189s
*/

func Sublist2(l1, l2 []int) Relation {

	if len(l1) == 0 && len(l2) == 0 {
		return RelationEqual
	}

	if len(l1) > 0 && len(l2) == 0 {
		return RelationSuperlist
	}

	if len(l1) == 0 && len(l2) > 0 {
		return RelationSublist
	}

	var s1, s2 string

	for _, n := range l1 {
		s1 += strconv.Itoa(n) + ","
	}

	for _, n := range l2 {
		s2 += strconv.Itoa(n) + ","
	}

	if s1 == s2 {
		return RelationEqual
	}

	if strings.Contains(s1, s2) {
		return RelationSuperlist
	}

	if strings.Contains(s2, s1) {
		return RelationSublist
	}

	return RelationUnequal

}
