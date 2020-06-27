package main

import (
	"flag"
	"log"

	"github.com/MishaNiki/lsait/backend/internal/app/servers/auth"
	"github.com/MishaNiki/lsait/backend/internal/app/utils"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "./configs/auth.json", "path to config")
}

func main() {
	flag.Parse()

	config := auth.NewConfig()
	if err := utils.DecodeJSONFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	server := auth.New()
	if err := server.Configure(config); err != nil {
		log.Fatal(err)
	}

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
