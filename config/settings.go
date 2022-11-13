package config

import "github.com/go-play/utils"

var DB_USERNAME = utils.GetEnv("DB_USERNAME", "root")
var DB_PASSWORD = utils.GetEnv("DB_PASSWORD", "root")
var DB_NAME = utils.GetEnv("DB_NAME", "gotest")
var DB_HOST = utils.GetEnv("DB_HOST", "database")
var DB_PORT = utils.GetEnv("DB_PORT", "3306")
var PORT = utils.GetEnv("PORT", "8080")
