package config

import (
	"encoding/json"
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
	var defult_cfg = NewAppConfig()
	var defult_use_flag = false
	//unmarshal defult config from defult_config.json
	file, err := os.ReadFile("./defult_config.json")

	json.Unmarshal(file, defult_cfg)
	var env_var string
	var ok bool
	//set grpc env var
	env_var, ok = os.LookupEnv("grpc_address")
	if !ok {
		a.GrpcConfig.Address = defult_cfg.GrpcConfig.Address
		defult_use_flag = true
	} else if env_var == "" {
		return errors.New("empty env var input")
	}

	env_var, ok = os.LookupEnv("grpc_network")
	if !ok {
		a.GrpcConfig.Network = defult_cfg.GrpcConfig.Network
		defult_use_flag = true
	} else if env_var == "" {
		return errors.New("empty env var input")
	}
	//set database env var
	//set db_user
	env_var, ok = os.LookupEnv("db_user")
	if !ok {
		a.DbConfig.User = defult_cfg.DbConfig.User
		defult_use_flag = true
	} else if env_var == "" {
		return errors.New("empty env var input")
	}
	//set db_password
	env_var, ok = os.LookupEnv("db_password")
	if !ok {
		a.DbConfig.Password = defult_cfg.DbConfig.Password
		defult_use_flag = true
	} else if env_var == "" {
		return errors.New("empty env var input")
	}
	//set db_sslmode
	env_var, ok = os.LookupEnv("db_sslmode")
	if !ok {
		a.DbConfig.Sslmode = defult_cfg.DbConfig.Sslmode
		defult_use_flag = true
	} else if env_var == "" {
		return errors.New("empty env var input")
	}
	//set db_host
	env_var, ok = os.LookupEnv("db_host")
	if !ok {
		a.DbConfig.Host = defult_cfg.DbConfig.Host
		defult_use_flag = true
	} else if env_var == "" {
		return errors.New("empty env var input")
	}
	//set db_port
	env_var, ok = os.LookupEnv("db_port")
	if !ok {
		a.DbConfig.Port = defult_cfg.DbConfig.Port
		defult_use_flag = true
	} else if env_var == "" {
		return errors.New("empty env var input")
	}
	//set db_name
	env_var, ok = os.LookupEnv("db_name")
	if !ok {
		a.DbConfig.DbName = defult_cfg.DbConfig.DbName
		defult_use_flag = true
	} else if env_var == "" {
		return errors.New("empty env var input")
	}

	if defult_use_flag && err != nil {
		return err
	}
	return nil
}
