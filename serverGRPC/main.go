package main

import (
	"log"

	"github.com/Ali-Farhadnia/serverGRPC/bookserver"
	"github.com/Ali-Farhadnia/serverGRPC/cmd"
	"github.com/Ali-Farhadnia/serverGRPC/models/book"
)

func main() {
	//set log flag
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := cmd.Config()
	if err != nil {
		panic(err)
	}
	err = book.CreateBooksTable()
	if err != nil {
		panic(err)
	}
	log.Println("server is runing...")
	bookserver.Start()
}
