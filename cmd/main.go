package main

import (
	"fmt"
	"log"

	"github.com/QuocAnh189/DemoEcom/cmd/api"
)

func main() {
	server := api.NewAPIServer(fmt.Sprintf(":%s", "9090"))
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}