package commands

import "redis-like/internal/db"

func ZAddCommand(rdb *db.RedisLikeDB, key string, score string, member string) {
	rdb.SortedSet.ZAdd(key, score, member)
}
