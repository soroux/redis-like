package commands

import "redis-like/internal/db"

func ZRankCommand(rdb *db.RedisLikeDB, key string, member string) (int, bool) {
	sl := rdb.GetOrCreateSkipList(key)
	return sl.ZRank(member)
}
