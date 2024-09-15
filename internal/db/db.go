package db

import "sync"

type RedisLikeDB struct {
	Keys       *RadixTree
	HashTable  *HashTable
	SortedSets map[string]*SkipList
	mu         sync.RWMutex
}

func NewRedisLikeDB() *RedisLikeDB {
	return &RedisLikeDB{
		Keys:       NewRadixTree(),
		HashTable:  NewHashTable(),
		SortedSets: make(map[string]*SkipList),
	}
}

func (rdb *RedisLikeDB) GetOrCreateSkipList(key string) *SkipList {
	rdb.mu.Lock()
	defer rdb.mu.Unlock()

	if sl, exists := rdb.SortedSets[key]; exists {
		return sl
	}
	sl := NewSkipList()
	rdb.SortedSets[key] = sl
	return sl
}
