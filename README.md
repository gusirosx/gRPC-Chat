# gRPC Guide
Based on this [tutorial](https://www.youtube.com/watch?v=BdzYdN_Zd9Q)

## How to run this example

1. run the grpc server

```sh
$ go run server/server.go
```
or
```sh
$ make run_server
```
2. run the client

```sh
$ go run client/client.go
```
or
```sh
$ make run_client
```

## How to create proto files

1. use the makefile

```sh
$ make generate