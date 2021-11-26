package connections

import (
	"database/sql"

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
