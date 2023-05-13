package config

import "time"

const (
	MYSQL_HOST     = "docker"
	MYSQL_PORT     = 3309
	MYSQL_USERNAME = "pogf"
	MYSQL_PASSWORD = "123456"
	MYSQL_DATABASE = "xaut"
)

// 定义过期时间
const ExpireDuration = time.Hour * 24

// 定义 JWT 的 key
var JwtKey = []byte("POGF")
