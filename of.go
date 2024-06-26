package slice

import "github.com/binaryphile/iterator"

type OfBool = Of[bool]
type OfByte = Of[byte]
type OfFloat32 = Of[float32]
type OfFloat64 = Of[float64]
type OfInt = Of[int]
type OfInt32 = Of[int32]
type OfInt64 = Of[int64]
type OfRune = Of[rune]
type OfString = Of[string]

type Of[T comparable] []T

func (ts Of[T]) Contains(t T) bool {
	return ts.Index(t) != -1
}

// Filter returns the slice of elements from ts for which fn returns true.
func (ts Of[T]) Filter(fn func(T) bool) Of[T] {
	results := make([]T, 0, len(ts))

	for _, t := range ts {
		if fn(t) {
			results = append(results, t)
		}
	}

	return results
}

// ForEach applies fn, which has no return value, to each member of ts.
func (ts Of[T]) ForEach(fn func(T)) {
	for _, t := range ts {
		fn(t)
	}
}

func (ts Of[T]) Index(t T) int {
	for i := range ts {
		if t == ts[i] {
			return i
		}
	}

	return -1
}

func (ts Of[T]) Iterator() iterator.Iterator[T] {
	return iterator.FromSlice(ts)
}

// Map returns the slice resulting from applying fn, whose return type is the same as the elements of ts, to each member of ts.
func (ts Of[T]) Map(fn func(T) T) Of[T] {
	results := make([]T, len(ts))

	for i := range ts {
		results[i] = fn(ts[i])
	}

	return results
}

// MapToBool returns the slice resulting from applying fn, whose return type is bool, to each member of ts.
func (ts Of[T]) MapToBool(fn func(T) bool) OfBool {
	results := make([]bool, len(ts))

	for i := range ts {
		results[i] = fn(ts[i])
	}

	return results
}

// MapToInt returns the slice resulting from applying fn, whose return type is int, to each member of ts.
func (ts Of[T]) MapToInt(fn func(T) int) OfInt {
	results := make([]int, len(ts))

	for i := range ts {
		results[i] = fn(ts[i])
	}

	return results
}

// MapToStrSlice returns the slice resulting from applying fn, whose return type is []string, to each member of ts.
func (ts Of[T]) MapToStrSlice(fn func(T) []string) OfStrSlice {
	results := make([][]string, len(ts))

	for i := range ts {
		results[i] = fn(ts[i])
	}

	return results
}

// MapToStr returns the slice resulting from applying fn, whose return type is string, to each member of ts.
func (ts Of[T]) MapToStr(fn func(T) string) OfString {
	results := make([]string, len(ts))

	for i := range ts {
		results[i] = fn(ts[i])
	}

	return results
}
