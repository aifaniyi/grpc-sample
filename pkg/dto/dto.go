package dto

//go:generate sh -c "rm *.pb.go ; protoc --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative --proto_path=. *.proto"
