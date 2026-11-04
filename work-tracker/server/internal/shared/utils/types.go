package utils

func Deref[T any](nullable *T, def T) T {
	if nullable == nil {
		return def
	}
	return *nullable
}