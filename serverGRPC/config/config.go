package config

import "errors"

type grpc_config struct {
	Network string
	Address string
}
type database_config struct {
	User     string
	Password string
	Sslmode  string
	Host     string
	Port     string
	DbName   string
}
type AppConfig struct {
	GrpcConfig grpc_config
	DbConfig   database_config
}
type AppConfigger struct {
	app    *AppConfig
	errors []error
}
type GrpcConfigger struct {
	AppConfigger
}
type DbConfigger struct {
	AppConfigger
}

func NewAppCongigger() *AppConfigger {
	return &AppConfigger{&AppConfig{}, make([]error, 0)}
}
func (b *AppConfigger) DataBaseConfig() *DbConfigger {
	return &DbConfigger{*b}
}
func (d *DbConfigger) User(user string) *DbConfigger {
	if user == "" {
		d.errors = append(d.errors, errors.New("please set database user"))
		return d
	}
	d.app.DbConfig.User = user
	return d
}
func (d *DbConfigger) Password(password string) *DbConfigger {
	if password == "" {
		d.errors = append(d.errors, errors.New("please set database password"))
		return d
	}
	d.app.DbConfig.Password = password
	return d
}
func (d *DbConfigger) Sslmode(sslmode string) *DbConfigger {
	if sslmode == "" {
		d.errors = append(d.errors, errors.New("please set database sslmode"))
		return d
	}
	d.app.DbConfig.Sslmode = sslmode
	return d
}
func (d *DbConfigger) Host(host string) *DbConfigger {
	if host == "" {
		d.errors = append(d.errors, errors.New("please set database host"))
		return d
	}
	d.app.DbConfig.Host = host
	return d
}
func (d *DbConfigger) Port(port string) *DbConfigger {
	if port == "" {
		d.errors = append(d.errors, errors.New("please set database port"))
		return d
	}
	d.app.DbConfig.Port = port
	return d
}
func (d *DbConfigger) DbName(dbname string) *DbConfigger {
	if dbname == "" {
		d.errors = append(d.errors, errors.New("please set database name"))
		return d
	}
	d.app.DbConfig.DbName = dbname
	return d
}

//--------------------------------

func (b *AppConfigger) GrpcConfig() *GrpcConfigger {
	return &GrpcConfigger{*b}
}
func (g *GrpcConfigger) Network(network string) *GrpcConfigger {
	if network == "" {
		g.errors = append(g.errors, errors.New("please set grpc network"))
		return g
	}
	g.app.GrpcConfig.Network = network
	return g
}
func (g *GrpcConfigger) Address(address string) *GrpcConfigger {
	if address == "" {
		g.errors = append(g.errors, errors.New("please set grpc addres"))
		return g
	}
	g.app.GrpcConfig.Address = address
	return g
}

func (a *AppConfigger) ConfigApp() (*AppConfig, []error) {
	if len(a.errors) != 0 {
		return nil, a.errors
	}
	return a.app, nil
}
