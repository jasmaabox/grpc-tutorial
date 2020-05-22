# gRPC Tutorial Follow-Along

Following [TutorialEdge gRPC tutorial](https://www.youtube.com/watch?v=BdzYdN_Zd9Q)

## Setup

  - Download [protoc](https://github.com/protocolbuffers/protobuf/releases/) and add to `PATH`
  - Install protoc generator for Go (this one supports plugins): `go install github.com/golang/protobuf/protoc-gen-go`

### Generate gRPC code
    protoc --go_out=plugins=grpc:chat chat.proto

### Run
    # Run in separate terminals
    go run server.go
    go run client. go

## Notes

### gRPC
  - Remote procedure call framework
  - Used in distributed system to expose methods
    in one application for use in another
  - Similar to REST

### Compared to REST
  - gRPC uses HTTP/2, REST uses HTTP/1.1
  - gRPC uses protocol buffer data format -> better performance versus JSON
  - Harder to use (cant really use with REST tools like Postman)
