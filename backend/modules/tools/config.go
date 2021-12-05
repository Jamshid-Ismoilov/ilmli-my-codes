package tools

import "github.com/go-redis/redis"

var Client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

var GSender = ""
var GPassword = ""
var GHost = "smtp.gmail.com"
var GPort = "587"
