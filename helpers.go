package vk

// ElemInSlice checks if element is in the slice
func ElemInSlice(elem string, slice []string) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}
