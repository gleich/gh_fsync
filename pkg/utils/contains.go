package utils

// Check if a string is a key in a map
func ContainsMapKey(k string, m map[string]string) bool {
	for key := range m {
		if key == k {
			return true
		}
	}
	return false
}
