package commands

import "redis-like/internal/db"

func ZAddCommand(rdb *db.RedisLikeDB, key string, score string, member string) {
	sl := rdb.GetOrCreateSkipList(key)
	sl.ZAdd(score, member)
}
