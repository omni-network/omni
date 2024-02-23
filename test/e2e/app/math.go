package app

func sum(batches []int) uint64 {
	var resp int
	for _, b := range batches {
		resp += b
	}

	return uint64(resp)
}
