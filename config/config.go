package config

import (
	"strconv"

	"github.com/mozhario/go_talk/utils"
)

var ServerHost string = utils.GetEnv("SERVER_HOST", "localhost")
var ServerPort string = utils.GetEnv("SERVER_HOST", "8080")
var WebscketPort string = utils.GetEnv("SERVER_HOST", "8090")

var ServerBufferSize, _ = strconv.Atoi(utils.GetEnv("SERVER_HOST", "1024"))
