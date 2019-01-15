package main

import (
	"gin-micro/routers/handlers"
	"github.com/micro/go-web"
	"log"
)

func main() {
	ws := web.NewService(web.Name("web"), web.Address(":20050"))

	r := handlers.InitRouter()

	ws.Handle("/", r)

	if err := ws.Init(); err != nil {
		log.Fatal("init--->", err)
	}

	if err := ws.Run(); err != nil {
		log.Fatalf("run err %v", err)
	}
}
