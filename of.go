package slice

type OfBool = Of[bool]
type OfInt = Of[int]
type OfString = Of[string]
type OfStrSlice = Of[[]string]

type Of[T any] []T

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