package commands

import "redis-like/internal/db"

func HSetCommand(rdb *db.RedisLikeDB, key string, field string, value interface{}) {
	rdb.HashTable.HSet(key, field, value)
}
