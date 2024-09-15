package commands

import "redis-like/internal/db"

func ZRankCommand(rdb *db.RedisLikeDB, key string, member string) (int, bool) {
	return rdb.SortedSet.ZRank(key, member)
}
