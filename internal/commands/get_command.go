package commands

import "redis-like/internal/db"

func GetCommand(rdb *db.RedisLikeDB, key string) (interface{}, bool) {
	return rdb.Keys.Get(key)
}
