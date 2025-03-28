package app

func sum(batches []int) uint64 {
	var resp int
	for _, b := range batches {
		resp += b
	}

	if resp < 0 {
		return 0
	}

	return uint64(resp)
}
