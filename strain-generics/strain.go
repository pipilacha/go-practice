package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

/*
Generics allows to declare a fuction that we can re-use for data of different types
func name[T type]	->  this tells go any param or return param of type T is of the specified type

func Keep[T any](elements []T, predicate func(T) bool) (kept []T) {
	for _, v := range elements {
		if predicate(v) {
			kept = append(kept, v)
		}

	}
	return
}

In the example above the function Keep takes a slice of elements of any type, a function
that takes a param of any type and returns another slice of elements of any type.
*/

// We can use an interface to allow only data of specific types
type Slicer interface {
	int | string | []int
}

// T is of type Slicer which can be of type int, string or []int
func Keep[T Slicer](elements []T, predicate func(T) bool) (kept []T) {
	for _, v := range elements {
		if predicate(v) {
			kept = append(kept, v)
		}

	}
	return
}

func Discard[T Slicer](elements []T, predicate func(T) bool) (discarded []T) {
	for _, v := range elements {
		if !predicate(v) {
			discarded = append(discarded, v)
		}
	}
	return
}
