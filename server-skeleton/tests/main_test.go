package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pressly/goose/v3"
	"os"
	"server-skeleton/api_init"
	"testing"
)

var Router *gin.Engine

func TestMain(m *testing.M) {
	fmt.Println("Init tests ...")
	err := api_init.MainInit("../")
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		api_init.InitGlobal.Cfg.DbHost,
		api_init.InitGlobal.Cfg.DbUser,
		api_init.InitGlobal.Cfg.DbPassword,
		api_init.InitGlobal.Cfg.DbName,
		api_init.InitGlobal.Cfg.DbPort,
		api_init.InitGlobal.Cfg.DbSSLMode,
	)

	db, err := goose.OpenDBWithDriver(api_init.InitGlobal.Cfg.DbDriver, dsn)

	if err != nil {
		panic(err)
	}
	err = goose.Up(db, "./../db/migrations")
	if err != nil {
		panic(err)
	}

	Router = gin.Default()

	os.Exit(m.Run())
}
