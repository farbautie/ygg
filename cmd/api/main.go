package main

import (
	"github.com/farbautie/ygg/config"
	api "github.com/farbautie/ygg/internal/app"
)

func main() {
	cfg := config.NewConfig()
	api.Run(cfg)
}
