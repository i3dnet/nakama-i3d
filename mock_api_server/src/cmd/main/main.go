package main

import (
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/mock-api-server/api"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/mock-api-server/internal/configs"
	"log"
)

func main() {

	config := configs.NewConfig()
	mapper := api.NewGinMapping(config)

	err := mapper.Run()

	if err != nil {
		log.Fatalln(err)
	}
}
