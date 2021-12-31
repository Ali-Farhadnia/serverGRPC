# Grpc server sample

A simple server that models a book repository where book information can be stored.

## book

```go
package book

type Book struct {
	ID        string `json="id"`
	Name      string `json="name"`
	Author    string `json="author"`
	Pagecount int32  `json="pagecount"`
	Inventory int64  `json="inventory"`
}
```
## env vars
grpc_address (need to specified)

grpc_network (default="tcp")

db_user (database user)(default="postgres")

db_password (database password)(need to specified)

db_sslmode (database sslmode)(default="disable")

db_host (database host)(default="localhost")

db_port (database port)(default="5432")

db_dbname (database name )(default="book")

## some points

This program uses the postgresql database and expects to already have a database with the name specified in the env vars(db_dbname) section.

This program uses GRPC for communication and the used GRPC model can be found in the ./models/modelpb/model.proto .
I also wrote a client program to use this program, which you can access from this [link](https://github.com/Ali-Farhadnia/sclientGRPC).

You can also run the program using Docker Compose.
