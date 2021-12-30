package config_test

import (
	"testing"

	"github.com/Ali-Farhadnia/serverGRPC/config"
	"github.com/stretchr/testify/assert"
)

func TestSetEnvVar(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		app := config.NewAppConfig()
		t.Setenv("grpc_address", "grpc_address")
		t.Setenv("grpc_network", "grpc_network")
		t.Setenv("db_user", "db_user")
		t.Setenv("db_password", "db_password")
		t.Setenv("db_sslmode", "db_sslmode")
		t.Setenv("db_host", "db_host")
		t.Setenv("db_port", "db_port")
		t.Setenv("db_dbname", "db_dbname")
		err := app.SetEnvVar()

		assert.NoError(t, err)
		if app.GrpcConfig.Address != "grpc_address" || app.GrpcConfig.Network != "grpc_network" ||
			app.DbConfig.User != "db_user" || app.DbConfig.Password != "db_password" ||
			app.DbConfig.Sslmode != "db_sslmode" || app.DbConfig.Host != "db_host" ||
			app.DbConfig.Port != "db_port" || app.DbConfig.DbName != "db_dbname" {
			t.Error("set all value failed")

		}
	})
	// errors.
	t.Run("UnsetGrpcAddress", func(t *testing.T) {
		app := config.NewAppConfig()
		t.Setenv("grpc_network", "grpc_network")
		t.Setenv("db_user", "db_user")
		t.Setenv("db_password", "db_password")
		t.Setenv("db_sslmode", "db_sslmode")
		t.Setenv("db_host", "db_host")
		t.Setenv("db_port", "db_port")
		t.Setenv("db_dbname", "db_dbname")
		err := app.SetEnvVar()
		assert.EqualError(t, err, "empty env var input:grpc_address")
	})

	t.Run("UnsetDbPassword", func(t *testing.T) {
		app := config.NewAppConfig()
		t.Setenv("grpc_address", "grpc_address")
		t.Setenv("grpc_network", "grpc_network")
		t.Setenv("db_user", "db_user")
		t.Setenv("db_sslmode", "db_sslmode")
		t.Setenv("db_host", "db_host")
		t.Setenv("db_port", "db_port")
		t.Setenv("db_dbname", "db_dbname")
		err := app.SetEnvVar()

		assert.EqualError(t, err, "empty env var input:db_password")
	})
	// defaults.

	t.Run("DefaultGrpcNetwork", func(t *testing.T) {
		app := config.NewAppConfig()
		t.Setenv("grpc_address", "grpc_address")
		t.Setenv("db_user", "db_user")
		t.Setenv("db_password", "db_password")
		t.Setenv("db_sslmode", "db_sslmode")
		t.Setenv("db_host", "db_host")
		t.Setenv("db_port", "db_port")
		t.Setenv("db_dbname", "db_dbname")
		app.SetEnvVar()

		assert.Equal(t, app.GrpcConfig.Network, "tcp")
	})

	t.Run("DefaultDbUser", func(t *testing.T) {
		app := config.NewAppConfig()
		t.Setenv("grpc_address", "grpc_address")
		t.Setenv("grpc_network", "grpc_network")
		t.Setenv("db_password", "db_password")
		t.Setenv("db_sslmode", "db_sslmode")
		t.Setenv("db_host", "db_host")
		t.Setenv("db_port", "db_port")
		t.Setenv("db_dbname", "db_dbname")
		app.SetEnvVar()

		assert.Equal(t, app.DbConfig.User, "postgres")
	})
	t.Run("DefaultDbSslmode", func(t *testing.T) {
		app := config.NewAppConfig()
		t.Setenv("grpc_address", "grpc_address")
		t.Setenv("grpc_network", "grpc_network")
		t.Setenv("db_user", "db_user")
		t.Setenv("db_password", "db_password")
		t.Setenv("db_host", "db_host")
		t.Setenv("db_port", "db_port")
		t.Setenv("db_dbname", "db_dbname")
		app.SetEnvVar()

		assert.Equal(t, app.DbConfig.Sslmode, "disable")
	})
	t.Run("DefaultDbhHost", func(t *testing.T) {
		app := config.NewAppConfig()
		t.Setenv("grpc_address", "grpc_address")
		t.Setenv("grpc_network", "grpc_network")
		t.Setenv("db_user", "db_user")
		t.Setenv("db_password", "db_password")
		t.Setenv("db_sslmode", "db_sslmode")
		t.Setenv("db_port", "db_port")
		t.Setenv("db_dbname", "db_dbname")
		app.SetEnvVar()

		assert.Equal(t, app.DbConfig.Host, "localhost")
	})

	t.Run("DefaultDbPort", func(t *testing.T) {
		app := config.NewAppConfig()
		t.Setenv("grpc_address", "grpc_address")
		t.Setenv("grpc_network", "grpc_network")
		t.Setenv("db_user", "db_user")
		t.Setenv("db_password", "db_password")
		t.Setenv("db_sslmode", "db_sslmode")
		t.Setenv("db_host", "db_host")
		t.Setenv("db_dbname", "db_dbname")
		app.SetEnvVar()

		assert.Equal(t, app.DbConfig.Port, "5432")
	})

	t.Run("DefaultDbDbName", func(t *testing.T) {
		app := config.NewAppConfig()
		t.Setenv("grpc_address", "grpc_address")
		t.Setenv("grpc_network", "grpc_network")
		t.Setenv("db_user", "db_user")
		t.Setenv("db_password", "db_password")
		t.Setenv("db_sslmode", "db_sslmode")
		t.Setenv("db_host", "db_host")
		t.Setenv("db_port", "db_port")
		app.SetEnvVar()

		assert.Equal(t, app.DbConfig.DbName, "book")
	})
}
