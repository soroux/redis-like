package main

import (
	"bufio"
	"fmt"
	"os"
	"redis-like/internal/commands"
	"redis-like/internal/db"
	"strings"
)

func main() {
	rdb := db.NewRedisLikeDB()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Redis-like CLI! Type 'exit' to quit.")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if strings.TrimSpace(input) == "exit" {
			break
		}
		processCommand(rdb, input)
	}
}

func processCommand(rdb *db.RedisLikeDB, command string) {
	parts := strings.SplitN(command, " ", 4)
	if len(parts) < 2 {
		fmt.Println("Invalid command")
		return
	}

	cmd, key := parts[0], parts[1]

	switch cmd {
	case "SET":
		if len(parts) < 3 {
			fmt.Println("Invalid SET command. Usage: SET key value")
			return
		}
		value := parts[2]
		commands.SetCommand(rdb, key, value)
		fmt.Printf("Key '%s' set to '%s'\n", key, value)

	case "GET":
		value, exists := commands.GetCommand(rdb, key)
		if exists {
			fmt.Printf("Value for '%s': %s\n", key, value)
		} else {
			fmt.Printf("Key '%s' not found\n", key)
		}

	case "HSET":
		if len(parts) < 4 {
			fmt.Println("Invalid HSET command. Usage: HSET key field value")
			return
		}
		field, value := parts[2], parts[3]
		commands.HSetCommand(rdb, key, field, value)

	case "HGET":
		if len(parts) < 3 {
			fmt.Println("Invalid HGET command. Usage: HGET key field")
			return
		}
		field := parts[2]
		value, exists := commands.HGetCommand(rdb, key, field)
		if exists {
			fmt.Printf("Value for field '%s' in hash '%s': %s\n", field, key, value)
		} else {
			fmt.Printf("Field '%s' not found in hash '%s'\n", field, key)
		}

	case "ZADD":
		if len(parts) < 4 {
			fmt.Println("Invalid ZADD command. Usage: ZADD key score member")
			return
		}
		score := parts[2]
		member := parts[3]
		// Convert score to float64
		// Ensure error handling
		commands.ZAddCommand(rdb, key, score, member)
		fmt.Printf("Member '%s' with score '%s' added to sorted set '%s'\n", member, score, key)

	case "ZRANGE":
		members := commands.ZRangeCommand(rdb, key)
		fmt.Printf("Members in sorted set '%s': %v\n", key, members)

	case "ZSCORE":
		if len(parts) < 3 {
			fmt.Println("Invalid ZSCORE command. Usage: ZSCORE key member")
			return
		}
		member := parts[2]
		score, exists := commands.ZScoreCommand(rdb, key, member)
		if exists {
			fmt.Printf("Score for member '%s' in sorted set '%s': %f\n", member, key, score)
		} else {
			fmt.Printf("Member '%s' not found in sorted set '%s'\n", member, key)
		}

	case "ZRANK":
		if len(parts) < 3 {
			fmt.Println("Invalid ZRANK command. Usage: ZRANK key member")
			return
		}
		member := parts[2]
		rank, exists := commands.ZRankCommand(rdb, key, member)
		if exists {
			fmt.Printf("Rank for member '%s' in sorted set '%s': %d\n", member, key, rank)
		} else {
			fmt.Printf("Member '%s' not found in sorted set '%s'\n", member, key)
		}

	default:
		fmt.Println("Unknown command:", cmd)
	}
}
