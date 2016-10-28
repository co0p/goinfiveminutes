package episode0

type MemHashTable struct {
	entries map[string][]byte
}

func NewMemHashTable() HashTable {
	return &MemHashTable{entries: make(map[string][]byte)}
}

func (mht *MemHashTable) Get(key string) ([]byte, error) {
	val, ok := mht.entries[key]
	if !ok {
		return nil, ErrNotFound
	}

	return val, nil
}

func (mht *MemHashTable) Set(key string, value []byte) error {
	mht.entries[key] = value
	return nil
}
