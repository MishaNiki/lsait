package main

import (
	"flag"
	"log"

	"github.com/MishaNiki/lsait/backend/internal/app/servers/article"
	"github.com/MishaNiki/lsait/backend/internal/app/utils"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "./configs/article.json", "path to config")
}

func main() {
	flag.Parse()

	config := article.NewConfig()
	if err := utils.DecodeJSONFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	server := article.New()
	if err := server.Configure(config); err != nil {
		log.Fatal(err)
	}

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
