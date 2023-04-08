package mini

// Exists checks if value exists in slice.
func Exists[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}
