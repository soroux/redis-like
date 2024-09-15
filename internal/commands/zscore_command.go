package commands

import "redis-like/internal/db"

func ZScoreCommand(rdb *db.RedisLikeDB, key string, member string) (string, bool) {
	return rdb.SortedSet.ZScore(key, member)
}
