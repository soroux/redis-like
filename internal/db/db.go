package db

type RedisLikeDB struct {
	Keys      *RadixTree
	HashTable *HashTable
	SortedSet *SkipList
}

func NewRedisLikeDB() *RedisLikeDB {
	return &RedisLikeDB{
		Keys:      NewRadixTree(),
		HashTable: NewHashTable(),
		SortedSet: NewSkipList(),
	}
}
