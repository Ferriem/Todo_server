package config

import (
	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

type Err string

const userList = "userList"
