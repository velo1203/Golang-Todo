package main

import (
	"log"
	"net/http"

	configs "studioj/boilerplate_go/configs"
	"studioj/boilerplate_go/internal/database"
	"studioj/boilerplate_go/internal/helpers"
	"studioj/boilerplate_go/internal/routers"

	"github.com/apex/gateway"
)

func main() {
	Init()
	Run()
}

func Init() {
	database.Init()
	routers.Init()
	database.AutoMigrate()
	log.Println(configs.PORT, " server started")
}

func Run() {
	if helpers.InLambda() {
		log.Fatal(gateway.ListenAndServe(configs.PORT, routers.Router))
	} else {
		log.Fatal(http.ListenAndServe(configs.PORT, routers.Router))
	}
}
