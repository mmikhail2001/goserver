package main

import (
	"log"

	"github.com/mmikhail2001/goserver/handler"
	"github.com/mmikhail2001/goserver/service"
	"github.com/mmikhail2001/goserver/storage"
)

func main() {
	store := storage.NewStorage()
	serv := service.NewService(store)
	handle := handler.NewHandler(serv)
	server := handler.NewServer("127.0.0.1:8080", handle.InitRoutes())
	if err := server.Run(); err != nil {
		log.Fatalf(err)
	}

}
