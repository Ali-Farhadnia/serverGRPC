package myserver

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Ali-Farhadnia/serverGRPC/config"
	"github.com/Ali-Farhadnia/serverGRPC/models/book"
	"github.com/Ali-Farhadnia/serverGRPC/models/modelpb"
	"google.golang.org/grpc"
)

type server struct {
	modelpb.UnimplementedCRUDServer
}

func (s *server) InsertBook(ctx context.Context, input *modelpb.Books) (*modelpb.Status, error) {
	logg := "InsertBook:"
	var books []book.Book
	var book book.Book
	for _, inputbook := range input.Books {
		book.Name = inputbook.Name
		book.Author = inputbook.Author
		book.Pagecount = inputbook.Pagescount
		book.Inventory = inputbook.Inventory
		books = append(books, book)
	}
	output := &modelpb.Status{}
	result, err := config.App.DB.InsertBooks(books)
	if err != nil {
		logg += "failed"
		output.Status = "no"
		output.Description = ""
	} else {
		logg += "success"
		output.Status = "ok"
		output.Description = result
	}
	log.Println(logg)
	return output, nil
}

func (s *server) FindBookById(ctx context.Context, input *modelpb.BookID) (*modelpb.FindResponse, error) {
	logg := "FindBookById:"

	output := modelpb.FindResponse{}
	stat := &modelpb.Status{}
	book := &modelpb.Book{}
	result, err := config.App.DB.FindBookById(input.Id)
	if err != nil {

		logg += "failed"
		stat.Status = "no"
		stat.Description = err.Error()

	} else {
		logg += "success"
		stat.Status = "ok"
		stat.Description = ""
		book.Id = result.ID
		book.Name = result.Name
		book.Author = result.Author
		book.Pagescount = result.Pagecount
		book.Inventory = result.Inventory
	}
	log.Println(logg)
	output.Book = book
	output.Status = stat
	return &output, nil
}

func (s *server) UpdateBook(ctx context.Context, input *modelpb.UpdateRequest) (*modelpb.Status, error) {
	logg := "UpdateBook:"
	var book book.Book
	var id string
	book.Name = input.Book.Name
	book.Author = input.Book.Author
	book.Pagecount = input.Book.Pagescount
	book.Inventory = input.Book.Inventory
	id = input.Id
	output := &modelpb.Status{}
	result, err := config.App.DB.UpdateBook(book, id)
	if err != nil {
		logg += "failed"
		output.Status = "no"
		output.Description = result

	} else {
		logg += "success"
		output.Status = "ok"
	}
	output.Description = "success"
	log.Println(logg)
	return output, nil
}

func (s *server) DeleteBook(ctx context.Context, input *modelpb.BookID) (*modelpb.Status, error) {
	logg := "DeleteBook:"

	var id string
	id = input.Id
	output := &modelpb.Status{}
	result, err := config.App.DB.DeleteBook(id)
	if err != nil {
		logg += "failed"
		output.Status = "no"
		output.Description = result

	} else {
		logg += "success"
		output.Status = "ok"
	}
	output.Description = "success"
	log.Println(logg)
	return output, nil

}

func Start() {
	log.Println("Start lestening...")
	lis, err := net.Listen(config.App.GrpcConfig.Network, config.App.GrpcConfig.Address)
	if err != nil {
		log.Fatalf("lestening was failed. error : %v", err)
	}

	s := grpc.NewServer()
	modelpb.RegisterCRUDServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve :%v", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	lis.Close()

	log.Println("stop lestening")

}
