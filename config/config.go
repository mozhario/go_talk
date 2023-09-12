package config

import (
	"strconv"

	"github.com/mozhario/go_talk/utils"
)

var ServerHost string = utils.GetEnv("SERVER_HOST", "localhost")
var ServerPort string = utils.GetEnv("SERVER_HOST", "8080")
var WebscketPort string = utils.GetEnv("SERVER_HOST", "8090")

var ServerBufferSize, _ = strconv.Atoi(utils.GetEnv("SERVER_HOST", "1024"))

var DBName = utils.GetEnv("DB_NAME", "go_talk")
var DBUsername = utils.GetEnv("DB_USERNAME", "postgres")
var DBPassword = utils.GetEnv("DB_PASSWORD", "postgres")
var DBHost = utils.GetEnv("DB_HOST", "localhost")
var DBPort = utils.GetEnv("DB_PORT", "5432")
