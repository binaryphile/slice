package slice

import "github.com/binaryphile/iterator"

type OfTo[T, R comparable] []T

func (ts OfTo[T, R]) Contains(t T) bool {
	return ts.Index(t) != -1
}

// Filter returns the slice of elements from ts for which fn returns true.
func (ts OfTo[T, R]) Filter(fn func(T) bool) OfTo[T, R] {
	results := make([]T, 0, len(ts))

	for _, t := range ts {
		if fn(t) {
			results = append(results, t)
		}
	}

	return results
}

// ForEach applies fn, which has no return value, to each member of ts.
func (ts OfTo[T, R]) ForEach(fn func(T)) {
	for _, t := range ts {
		fn(t)
	}
}

func (ts OfTo[T, R]) Index(t T) int {
	for i := range ts {
		if t == ts[i] {
			return i
		}
	}

	return -1
}

func (ts OfTo[T, R]) Iterator() iterator.Iterator[T] {
	return iterator.FromSlice(ts)
}

func (ts OfTo[T, R]) Map(fn func(T) R) Of[R] {
	results := make([]R, len(ts))

	for i, t := range ts {
		results[i] = fn(t)
	}

	return results
}

// MapToBool returns the slice resulting from applying fn, whose return type is bool, to each member of ts.
func (ts OfTo[T, R]) MapToBool(fn func(T) bool) OfTo[bool, R] {
	results := make([]bool, len(ts))

	for i := range ts {
		results[i] = fn(ts[i])
	}

	return results
}

// MapToInt returns the slice resulting from applying fn, whose return type is int, to each member of ts.
func (ts OfTo[T, R]) MapToInt(fn func(T) int) OfTo[int, R] {
	results := make([]int, len(ts))

	for i := range ts {
		results[i] = fn(ts[i])
	}

	return results
}

// MapToStrSlice returns the slice resulting from applying fn, whose return type is []string, to each member of ts.
func (ts OfTo[T, R]) MapToStrSlice(fn func(T) []string) OfToAny[[]string, R] {
	results := make([][]string, len(ts))

	for i := range ts {
		results[i] = fn(ts[i])
	}

	return results
}

// MapToStr returns the slice resulting from applying fn, whose return type is string, to each member of ts.
func (ts OfTo[T, R]) MapToStr(fn func(T) string) OfTo[string, R] {
	results := make([]string, len(ts))

	for i := range ts {
		results[i] = fn(ts[i])
	}

	return results
}
