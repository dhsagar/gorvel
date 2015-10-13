package config

import (
	"github.com/robfig/config"
)

var configMap map[string]string

func ReadAllConfig() {
	cfg, _ := config.ReadDefault("config.cfg")
	configMap = make(map[string]string)

	configMap["app.env"], _ = cfg.String("CONFIG", "app.env")
	configMap["app.port"], _ = cfg.String("CONFIG", "app.port")


	configMap["db.type"], _ = cfg.String("DATABASE", "type")
	configMap["db.host"], _ = cfg.String("DATABASE", "url")
	configMap["db.port"], _ = cfg.String("DATABASE", "port")
	configMap["db.user"], _ = cfg.String("DATABASE", "username")
	configMap["db.password"], _ = cfg.String("DATABASE", "password")
	configMap["db.name"], _ = cfg.String("DATABASE", "database")
}
