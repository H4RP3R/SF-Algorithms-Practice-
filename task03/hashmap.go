// Реализуйте тип хеш-мап, который использует строки в качестве ключей и значений.
// Он должен реализовывать минимум три метода: Set, Get, Delete.

package hashmap

type hmItem struct {
	key   string
	value string
}

type hashmap struct {
	storage [][]hmItem
	size    uint64
}

func (h *hashmap) hash(key string) uint64 {
	var runeSum int32
	for _, r := range key {
		runeSum += r
	}

	return uint64(runeSum) % h.size
}

func (h *hashmap) Set(key, val string) {
	idx := h.hash(key)
	bucket := h.storage[idx]
	for i, item := range bucket {
		if item.key == key {
			bucket[i].value = val
			return
		}
	}
	h.storage[idx] = append(bucket, hmItem{key, val})
}

func (h *hashmap) Get(key string) (value string, ok bool) {
	idx := h.hash(key)
	bucket := h.storage[idx]
	for _, item := range bucket {
		if item.key == key {
			return item.value, true
		}
	}

	return "", false
}

func (h *hashmap) Delete(key string) {
	idx := h.hash(key)
	bucket := h.storage[idx]
	for i, item := range bucket {
		if item.key == key {
			h.storage[idx] = append(bucket[:i], bucket[i+1:]...)
			return
		}
	}
}

func NewHashmap(size uint64) *hashmap {
	return &hashmap{storage: make([][]hmItem, size), size: size}
}
