package commands

import "redis-like/internal/db"

func ZScoreCommand(rdb *db.RedisLikeDB, key string, member string) (string, bool) {
	sl := rdb.GetOrCreateSkipList(key)
	return sl.ZScore(member)
}
