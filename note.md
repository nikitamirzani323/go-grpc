https://github.com/codebangkok/golang
//SERVER
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get google.golang.org/grpc

//CLIENT
go get google.golang.org/grpc

//COMPILE PROTOC
protoc calculator.proto --go_out=../server --go-grpc_out=../server
protoc calculator.proto --go_out=../client --go-grpc_out=../client