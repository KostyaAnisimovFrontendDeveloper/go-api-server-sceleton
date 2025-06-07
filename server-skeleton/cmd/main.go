package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"server-skeleton/api/page/route"
	"server-skeleton/api_init"
)

func routes(config *api_init.InitGlobalStruct) error {
	r := gin.Default()

	route.InitPageRoutes(r)

	port := fmt.Sprintf(":%s", config.Cfg.ServerPort)
	err := r.Run(port)
	return err
}

func main() {

	fmt.Println("Init main ...")
	err := api_init.MainInit("")
	if err != nil {
		log.Fatal(err)
	}

	err = routes(api_init.InitGlobal)
	if err != nil {
		log.Fatal(err)
	}
}
