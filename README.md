# Redis-Like Application
## Overview

This is a simple Redis-like application implemented in Go. It provides basic functionalities for handling key-value pairs, hashes, and sorted sets, mimicking some of the core commands of Redis.

## Features

- **SET** and **GET** for basic key-value storage.
- **HSET** and **HGET** for hash maps.
- **ZADD**, **ZRANGE**, **ZSCORE**, and **ZRANK** for sorted sets.

## Installation

To get started with this application, follow these steps:

1. **Clone the Repository**

   `bash
   git clone https://github.com/soroux/redis-like.git
   `

2. **Navigate to the Project Directory**

   `bash
   cd redis-like
   `

3. **Install Dependencies**

   Ensure you have Go installed. Then run:

   `bash
   go mod tidy
   `

4. **Build the Application**

   `bash
   go build -o redis-like
   `

## Usage

You can run the application directly from the terminal. Start the application with:

`bash
./redis-like
`

You can then enter commands in the terminal to interact with the application.

## Commands

### Key-Value Store

- **SET**

  `bash
  SET key value
  `

  Sets the value for the given key.

- **GET**

  `bash
  GET key
  `

  Retrieves the value for the given key.

### Hashes

- **HSET**

  `bash
  HSET key field value
  `

  Sets the value for the specified field in the hash stored at key.

- **HGET**

  `bash
  HGET key field
  `

  Retrieves the value for the specified field in the hash stored at key.

### Sorted Sets

- **ZADD**

  `bash
  ZADD key score member
  `

  Adds a member to the sorted set stored at key with the specified score.

- **ZRANGE**

  `bash
  ZRANGE key
  `

  Retrieves all members in the sorted set stored at key.

- **ZSCORE**

  `bash
  ZSCORE key member
  `

  Retrieves the score of the specified member in the sorted set stored at key.

- **ZRANK**

  `bash
  ZRANK key member
  `

  Retrieves the rank of the specified member in the sorted set stored at key.

## Example

Here\'s an example of how to use the application:

```bash
# Start the application
./redis-like

# Set key-value pair
SET mykey "myvalue"

# Get value
GET mykey
# Output: myvalue

# Set hash field
HSET myhash field1 "value1"

# Get hash field
HGET myhash field1
# Output: value1

# Add to sorted set
ZADD myset 1.0 "member1"

# Range in sorted set
ZRANGE myset
# Output: [member1]

# Get score
ZSCORE myset member1
# Output: 1.0

# Get rank
ZRANK myset member1
# Output: 0
```

## Contributing

If you want to contribute to this project, feel free to submit issues, create pull requests, or offer suggestions.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.' > README.md
