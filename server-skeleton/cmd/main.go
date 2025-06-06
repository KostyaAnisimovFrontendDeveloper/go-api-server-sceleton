package main

import (
	"fmt"
	"server-skeleton/api/page/model"
	"server-skeleton/db"
	"server-skeleton/pkg/config"
)

func main() {

	fmt.Println("Init config ...")
	cfg := &config.Config{}
	cfg.InitConfig()

	fmt.Println("Env: ", cfg.Env)

	fmt.Println("Init db ...")

	dbh, err := db.ConnectionFactory(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Run migrations ...")
	err = dbh.AutoMigrate(&model.Page{})

	if err != nil {
		panic(err)
	}

	//	repo := repository.PostgresPageRepository(dbh)
	fmt.Println("End!")

}
