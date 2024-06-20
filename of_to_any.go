package slice

import "github.com/binaryphile/iterator"

type OfToAny[T, R any] []T

// Filter returns the slice of elements from ts for which fn returns true.
func (ts OfToAny[T, R]) Filter(fn func(T) bool) OfToAny[T, R] {
	results := make([]T, 0, len(ts))

	for _, t := range ts {
		if fn(t) {
			results = append(results, t)
		}
	}

	return results
}

// ForEach applies fn, which has no return value, to each member of ts.
func (ts OfToAny[T, R]) ForEach(fn func(T)) {
	for _, t := range ts {
		fn(t)
	}
}

func (ts OfToAny[T, R]) Iterator() iterator.Iterator[T] {
	return iterator.FromSlice(ts)
}

func (ts OfToAny[T, R]) Map(fn func(T) R) OfAny[R] {
	results := make([]R, len(ts))

	for i, t := range ts {
		results[i] = fn(t)
	}

	return results
}

// MapToBool returns the slice resulting from applying fn, whose return type is bool, to each member of ts.
func (ts OfToAny[T, R]) MapToBool(fn func(T) bool) OfToAny[bool, R] {
	results := make([]bool, len(ts))

	for i := range ts {
		results[i] = fn(ts[i])
	}

	return results
}

// MapToInt returns the slice resulting from applying fn, whose return type is int, to each member of ts.
func (ts OfToAny[T, R]) MapToInt(fn func(T) int) OfToAny[int, R] {
	results := make([]int, len(ts))

	for i := range ts {
		results[i] = fn(ts[i])
	}

	return results
}

// MapToStrSlice returns the slice resulting from applying fn, whose return type is []string, to each member of ts.
func (ts OfToAny[T, R]) MapToStrSlice(fn func(T) []string) OfToAny[[]string, R] {
	results := make([][]string, len(ts))

	for i := range ts {
		results[i] = fn(ts[i])
	}

	return results
}

// MapToStr returns the slice resulting from applying fn, whose return type is string, to each member of ts.
func (ts OfToAny[T, R]) MapToStr(fn func(T) string) OfToAny[string, R] {
	results := make([]string, len(ts))

	for i := range ts {
		results[i] = fn(ts[i])
	}

	return results
}
