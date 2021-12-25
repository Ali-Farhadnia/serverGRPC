package cmd

import (
	"github.com/Ali-Farhadnia/serverGRPC/config"
)

var AppConfig *config.AppConfig

func Config() []error {
	var errs = make([]error, 0)
	configger := config.NewAppCongigger()
	configger.GrpcConfig().
		Address("0.0.0.0:50051").
		Network("tcp").
		DataBaseConfig().
		User("postgres").
		Password("1234").
		Host("localhost").
		Port("5432").
		Sslmode("disable").
		DbName("book")
	AppConfig, errs = configger.ConfigApp()
	if len(errs) != 0 {
		return errs
	}
	return nil
}

/*
//set grpc configuration
func Configgrpc() error {
	config.App.GrpcConfig.Address = "0.0.0.0:50051"
	config.App.GrpcConfig.Network = "tcp"
	return nil
}

func ConfigDb() error {

	var err error
	//set database configuration
	config.App.DB.Config.Password = "1234"
	config.App.DB.Config.User = "postgres"
	config.App.DB.Config.Host = "localhost"
	config.App.DB.Config.Port = "5432"
	config.App.DB.Config.Sslmode = "disable"
	config.App.DB.Config.DbName = "book"

	//set client accses to book database
	err = config.App.DB.SetBookdb()
	if err != nil {
		log.Println(err)
	}
	//check if books table exist or not and if not exist create one
	err = config.App.DB.CreateBooksTable()
	if err != nil {
		log.Println(err)

	}
	return nil
}
*/
