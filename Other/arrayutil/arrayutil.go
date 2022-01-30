package arrayutil

// StringSliceContains returns true if element e is in slice s
func StringSliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
