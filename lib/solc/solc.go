package solc

type StorageLayout struct {
	Storage []StorageLayoutEntry         `json:"storage"`
	Types   map[string]StorageLayoutType `json:"types"`
}

type StorageLayoutEntry struct {
	AstID    uint   `json:"astId"`
	Contract string `json:"contract"`
	Label    string `json:"label"`
	Offset   uint   `json:"offset"`
	Slot     uint   `json:"slot,string"`
	Type     string `json:"type"`
}

type StorageLayoutType struct {
	Encoding      string `json:"encoding"`
	Label         string `json:"label"`
	NumberOfBytes uint   `json:"numberOfBytes,string"`
	Key           string `json:"key,omitempty"`
	Value         string `json:"value,omitempty"`
}

// SlotOf returns the slot number of the given label in the storage layout.
func SlotOf(layout StorageLayout, label string) (uint, bool) {
	for _, entry := range layout.Storage {
		if entry.Label == label {
			return entry.Slot, true
		}
	}

	return 0, false
}
