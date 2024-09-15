package utils

func FNV1a(key string) uint64 {
	const fnvOffsetBasis uint64 = 14695981039346656037
	const fnvPrime uint64 = 1099511628211

	hash := fnvOffsetBasis
	for i := 0; i < len(key); i++ {
		hash ^= uint64(key[i]) // Ensure we are converting byte to uint64
		hash *= fnvPrime
	}
	return hash
}
