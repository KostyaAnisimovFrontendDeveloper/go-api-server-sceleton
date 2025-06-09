package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"server-skeleton/api/page/route"
	"server-skeleton/api_init"
	_ "server-skeleton/docs"
)

func routes(config *api_init.InitGlobalStruct) error {
	r := gin.Default()

	route.InitPageRoutes(r)

	port := fmt.Sprintf(":%s", config.Cfg.ServerPort)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
