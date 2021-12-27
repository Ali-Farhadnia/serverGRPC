package connections

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/Ali-Farhadnia/serverGRPC/cmd"
	_ "github.com/lib/pq"
)

//SetClientAccessPoint - set postgree connection to work with
func GetDBClientAccessPoint(connStr string) (*sql.DB, error) {
	//connect to postgres
	res, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return res, nil
}

/* Used to create a singleton object of sql.DB client.
Initialized and exposed through GetBookDb().*/
var clientInstance *sql.DB

//Used during creation of singleton client object in GetBookDb().
var clientInstanceError error

//Used to execute client creation procedure only once.
var DbOnce sync.Once

//GetBookDb - set postgree connection to Book database to  work with
func GetBookDb() (*sql.DB, error) {
	DbOnce.Do(func() {
		dbconfig := cmd.AppConfig.DbConfig
		connStr := fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=%s",
			dbconfig.Host, dbconfig.Port, dbconfig.User, dbconfig.Password, dbconfig.DbName, dbconfig.Sslmode)
		res, err := GetDBClientAccessPoint(connStr)
		if err != nil {
			clientInstanceError = err
		}
		err = res.Ping()
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = res
	})

	return clientInstance, clientInstanceError
}
