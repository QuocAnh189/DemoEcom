package main

import (
	"fmt"
	"log"

	"github.com/QuocAnh189/DemoEcom/cmd/api"
	"github.com/QuocAnh189/DemoEcom/configs"
)

func main() {
	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port))
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}