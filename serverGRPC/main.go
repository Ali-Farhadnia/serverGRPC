package main

import (
	"log"

	"github.com/Ali-Farhadnia/serverGRPC/bookserver"
	"github.com/Ali-Farhadnia/serverGRPC/cmd"
)

func main() {
	//set log flag
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	errs := cmd.Config()
	if errs != nil {
		for _, v := range errs {
			log.Println(v)
		}
		panic("configuration failed")
	}
	/*
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
	*/
	log.Println("server is runing...")
	bookserver.Start()
}
