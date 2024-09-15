package db

import "redis-like/utils"

type HashTable struct {
	table map[uint64]map[string]interface{}
}

func NewHashTable() *HashTable {
	return &HashTable{table: make(map[uint64]map[string]interface{})}
}

func (h *HashTable) HSet(key string, field string, value interface{}) {
	hash := utils.FNV1a(key)
	if _, ok := h.table[hash]; !ok {
		h.table[hash] = make(map[string]interface{})
	}
	h.table[hash][field] = value
}

func (h *HashTable) HGet(key string, field string) (interface{}, bool) {
	hash := utils.FNV1a(key)
	if val, ok := h.table[hash][field]; ok {
		return val, true
	}
	return nil, false
}
