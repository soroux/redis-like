package commands

import "redis-like/internal/db"

func ZRangeCommand(rdb *db.RedisLikeDB, key string, start int, end int) []string {
	return rdb.SortedSet.ZRange(key, start, end)
}