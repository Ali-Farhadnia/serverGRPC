package config

import "github.com/Ali-Farhadnia/serverGRPC/models/bookdb"

type GrpcConfig struct {
	Network string
	Address string
}
type Appconfig struct {
	GrpcConfig GrpcConfig
	DB         bookdb.BookDB
}

var App Appconfig
