package main

import (
	"time"

	route "github.com/Prakashdurai/golang/api/route"
	"github.com/Prakashdurai/golang/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	if err := gin.Run(env.ServerAddress); err != nil {
    panic(err)
	}
}
