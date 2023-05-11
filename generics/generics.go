package generics

// valRef returns a pointer to the value passed regardless of type
func valRef[T any](v T) *T {
	return &v
}

// must takes a tuple return where the second return value is an error and coerces the return to a single value. If the
// error returned is not nil then the routine panics.
func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}
