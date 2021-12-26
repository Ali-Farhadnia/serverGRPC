package config

type grpc_config struct {
	Address string `json:"address"`
	Network string `json:"network"`
}
type database_config struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Sslmode  string `json:"sslmode"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"dbname"`
}
type AppConfig struct {
	GrpcConfig grpc_config     `json:"grpc_config"`
	DbConfig   database_config `json:"database_config"`
}

func NewAppConfig() *AppConfig {
	return &AppConfig{grpc_config{}, database_config{}}
}
