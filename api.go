package pasty

// handleIndex handles the current index of store.
func handleIndex(idx *int, size int) {
	*idx++
	if *idx >= size {
		*idx = 0
	}
}
