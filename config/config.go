package config

import (
	"errors"
	"os"
)

type GrpcConfig struct {
	Address string `json:"address"`
	Network string `json:"network"`
}
type DatabaseConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Sslmode  string `json:"sslmode"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"dbname"`
}
type AppConfig struct {
	GrpcConfig GrpcConfig     `json:"grpc_config"`
	DbConfig   DatabaseConfig `json:"database_config"`
}

func NewAppConfig() *AppConfig {
	return &AppConfig{GrpcConfig{}, DatabaseConfig{}}
}

func (a *AppConfig) SetEnvVar() error {
	var env_var string
	//set grpc env var
	env_var = os.Getenv("grpc_address")
	if env_var == "" {
		return errors.New("empty env var input:grpc_address")
	}
	a.GrpcConfig.Address = env_var
	env_var = os.Getenv("grpc_network")
	if env_var == "" {
		env_var = "tcp"
	}
	a.GrpcConfig.Network = env_var
	//set database env var
	//set db_user
	env_var = os.Getenv("db_user")
	if env_var == "" {
		env_var = "postgres"
	}
	a.DbConfig.User = env_var
	//set db_password
	env_var = os.Getenv("db_password")
	if env_var == "" {
		return errors.New("empty env var input:db_password")
	}
	a.DbConfig.Password = env_var
	//set db_sslmode
	env_var = os.Getenv("db_sslmode")
	if env_var == "" {
		env_var = "disable"
	}
	a.DbConfig.Sslmode = env_var
	//set db_host
	env_var = os.Getenv("db_host")
	if env_var == "" {
		env_var = "localhost"
	}
	a.DbConfig.Host = env_var
	//set db_port
	env_var = os.Getenv("db_port")
	if env_var == "" {
		env_var = "5432"
	}
	a.DbConfig.Port = env_var
	//set db_name
	env_var = os.Getenv("db_dbname")
	if env_var == "" {
		env_var = "book"
	}
	a.DbConfig.DbName = env_var

	return nil
}
