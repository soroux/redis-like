package commands

import "redis-like/internal/db"

func SetCommand(rdb *db.RedisLikeDB, key string, value interface{}) {
	rdb.Keys.Insert(key, value)
}
