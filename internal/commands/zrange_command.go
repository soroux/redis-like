package commands

import "redis-like/internal/db"

func ZRangeCommand(rdb *db.RedisLikeDB, key string) []string {
	sl := rdb.GetOrCreateSkipList(key)
	return sl.ZRange()
}
