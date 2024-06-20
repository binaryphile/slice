package slice

type OfBoolSlice = OfAny[[]bool]
type OfByteSlice = OfAny[[]byte]
type OfFloat32Slice = OfAny[[]float32]
type OfFloat64Slice = OfAny[[]float64]
type OfInt32Slice = OfAny[[]int32]
type OfInt64Slice = OfAny[[]int64]
type OfIntSlice = OfAny[[]int]
type OfRuneSlice = OfAny[[]rune]
type OfStrSlice = OfAny[[]string]

type OfAny[T any] []T

// Filter returns the slice of elements from ts for which fn returns true.
func (ts OfAny[T]) Filter(fn func(T) bool) OfAny[T] {
	results := make([]T, 0, len(ts))

	for _, t := range ts {
		if fn(t) {
			results = append(results, t)
		}
	}

	return results
}

// ForEach applies fn, which has no return value, to each member of ts.
func (ts OfAny[T]) ForEach(fn func(T)) {
	for _, t := range ts {
		fn(t)
	}
}

// Map returns the slice resulting from applying fn, whose return type is the same as the elements of ts, to each member of ts.
func (ts OfAny[T]) Map(fn func(T) T) OfAny[T] {
	results := make([]T, len(ts))

	for i := range ts {
		results[i] = fn(ts[i])
	}

	return results
}

// MapToBool returns the slice resulting from applying fn, whose return type is bool, to each member of ts.
func (ts OfAny[T]) MapToBool(fn func(T) bool) OfBool {
	results := make([]bool, len(ts))

	for i := range ts {
		results[i] = fn(ts[i])
	}

	return results
}

// MapToInt returns the slice resulting from applying fn, whose return type is int, to each member of ts.
func (ts OfAny[T]) MapToInt(fn func(T) int) OfInt {
	results := make([]int, len(ts))

	for i := range ts {
		results[i] = fn(ts[i])
	}

	return results
}

// MapToStrSlice returns the slice resulting from applying fn, whose return type is []string, to each member of ts.
func (ts OfAny[T]) MapToStrSlice(fn func(T) []string) OfStrSlice {
	results := make([][]string, len(ts))

	for i := range ts {
		results[i] = fn(ts[i])
	}

	return results
}

// MapToStr returns the slice resulting from applying fn, whose return type is string, to each member of ts.
func (ts OfAny[T]) MapToStr(fn func(T) string) OfString {
	results := make([]string, len(ts))

	for i := range ts {
		results[i] = fn(ts[i])
	}

	return results
}
