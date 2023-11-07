package main

import (
	"fmt"
	"os"

	"github.com/roby-aw/go-clean-architecture-hexagonal/app"
	"github.com/roby-aw/go-clean-architecture-hexagonal/config"
	"github.com/roby-aw/go-clean-architecture-hexagonal/utils"
)

// @title Go Clean Architecture Hexagonal API Documentation
// @description This is a sample server for a Go Clean Architecture Hexagonal API.
// @version 1.0.0
// @host localhost:8080
// @BasePath /v1
func main() {
	conf := config.GetConfig()
	dbCon := utils.NewConnectionDatabase(conf)
	server, port := app.Run(conf, dbCon)

	defer dbCon.CloseConnection()
	err := server.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	quit := make(chan os.Signal)
	<-quit
}
