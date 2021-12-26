package config

type AppConfig struct {
	//grpc
	GrpcAddress string ` long:"grpc_address" description:"grpc address" default:"0.0.0.0:50051"`
	GrpcNetwork string ` long:"grpc_network" description:"grpc network" default:"tcp"`
	//database
	DbUser     string `long:"db_user" description:"database user" default:"postgres"`
	DbPassword string `long:"db_password" description:"database password" default:"1234"`
	DbSslmode  string `long:"db_sslmode" description:"database ssl mode" default:"disable"`
	DbHost     string `long:"db_host" description:"database host" default:"localhost"`
	DbPort     string `long:"db_port" description:"database port" default:"5432"`
	DbDbName   string `long:"db_name" description:"database name" default:"book"`
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}
