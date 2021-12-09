# serverGRPC
A simple GRPC server that models a library.

The Postergs database is used to store information.

By using the [clientproject](https://github.com/Ali-Farhadnia/clientGRPC) , you can run the client server after running and perform the CRUD operation for the books.

You can also download the project Docker image from this [link](https://hub.docker.com/repository/docker/ali3242414268/server_grpc).

Note that the server is uses the port: 50051 to connect the GRPC

And uses the port: 5432 to connect to the database.

By default the database settings are as follows:

password = "1234"

user = "postgres"

host = "localhost"

port = "5432"

sslmode = "disable"

databasename = "book"

You can change this config in this path:  serverGRPC/cmd/setConfig
