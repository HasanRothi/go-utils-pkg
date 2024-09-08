package arrayutils

// SliceContains checks if a string is present in a slice of strings
func SliceContains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
