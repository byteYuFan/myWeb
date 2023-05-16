package config

import "time"

// 定义过期时间
const ExpireDuration = time.Hour * 24

// 定义 JWT 的 key
var JwtKey = []byte("POGF")
