package commands

import "redis-like/internal/db"

func HGetCommand(rdb *db.RedisLikeDB, key string, field string) (interface{}, bool) {
	return rdb.HashTable.HGet(key, field)
}
