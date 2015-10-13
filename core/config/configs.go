package config

import (
	"github.com/robfig/config"
	"fmt"
)

func ReadAllConfig() {
	cfg, _ := config.ReadDefault("config.cfg")

	fmt.Println(cfg.String("CONFIG", "mode"))
	fmt.Println(cfg.String("DATABASE", "url"))
}
