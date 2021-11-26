package main

import (
	"log"

	"github.com/Ali-Farhadnia/serverGRPC/cmd"
	"github.com/Ali-Farhadnia/serverGRPC/myserver"
)

func main() {
	//set log flag
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//config database and connect to it
	err := cmd.ConfigDb()
	if err != nil {
		panic(err)
	}

	err = cmd.Configgrpc()
	if err != nil {
		panic(err)
	}
	log.Println("server is runing...")
	myserver.Start()

}
