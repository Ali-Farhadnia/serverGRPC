package bookserver

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Ali-Farhadnia/serverGRPC/cmd"
	"github.com/Ali-Farhadnia/serverGRPC/models/book"
	"github.com/Ali-Farhadnia/serverGRPC/models/modelpb"
	"google.golang.org/grpc"
)

type server struct {
	modelpb.UnimplementedCRUDServer
}

func (s *server) InsertBook(ctx context.Context, input *modelpb.Books) (*modelpb.Status, error) {
	var logg = "InsertBook:"
	output := &modelpb.Status{}
	//var books []book.Book
	for _, inputbook := range input.Books {
		var book book.Book
		book.Name = inputbook.Name
		book.Author = inputbook.Author
		book.Pagecount = inputbook.Pagescount
		book.Inventory = inputbook.Inventory
		err := book.InsertToDb()
		if err != nil {
			logg += "failed"
			output.Status = "no"
			d, _ := book.String()
			output.Description = d
			return output, err
		}
		//books = append(books, book)
	}
	logg += "success"
	output.Status = "ok"
	output.Description = ""
	log.Println(logg)
	return output, nil
}

func (s *server) FindBookById(ctx context.Context, input *modelpb.BookID) (*modelpb.FindResponse, error) {
	logg := "FindBookById:"

	output := modelpb.FindResponse{}
	stat := &modelpb.Status{}
	b := &modelpb.Book{}
	result, err := book.FindBookById(input.Id)
	if err != nil {

		logg += "failed"
		stat.Status = "no"
		stat.Description = err.Error()

	} else {
		logg += "success"
		stat.Status = "ok"
		stat.Description = ""
		b.Id = result.ID
		b.Name = result.Name
		b.Author = result.Author
		b.Pagescount = result.Pagecount
		b.Inventory = result.Inventory
	}
	log.Println(logg)
	output.Book = b
	output.Status = stat
	return &output, nil
}

func (s *server) UpdateBook(ctx context.Context, input *modelpb.UpdateRequest) (*modelpb.Status, error) {
	logg := "UpdateBook:"
	var b book.Book
	var id string
	b.Name = input.Book.Name
	b.Author = input.Book.Author
	b.Pagecount = input.Book.Pagescount
	b.Inventory = input.Book.Inventory
	id = input.Id
	output := &modelpb.Status{}

	target, err := book.FindBookById(id)
	if err != nil {

		logg += "failed"
		output.Status = "no"
		output.Description = err.Error()

	}
	err = (*target).UpdateBook(&b)
	if err != nil {
		logg += "failed"
		output.Status = "no"
		output.Description = ""

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

	var id = input.Id
	output := &modelpb.Status{}
	err := book.DeleteBook(id)
	if err != nil {
		logg += "failed"
		output.Status = "no"
		output.Description = ""

	} else {
		logg += "success"
		output.Status = "ok"
	}
	output.Description = "success"
	log.Println(logg)
	return output, nil

}

func Start() {
	grpcconfig := cmd.AppConfig
	log.Println("Start lestening...")
	lis, err := net.Listen(grpcconfig.GrpcNetwork, grpcconfig.GrpcAddress)
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
