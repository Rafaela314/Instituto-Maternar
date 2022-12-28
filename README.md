## Protobuf Generation
### Golang
Go code must be generated while working as it is required for the application.

Please use protoc@v3.19.4. Installation instructions found [here](https://grpc.io/docs/protoc-installation/).  
To generate the Go files you must install `protoc-gen-go` and `protoc-gen-go-grpc`:
```console
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
```

Run:
```console
# protoc --proto_path=proto --go_out=$GOPATH/src --go-grpc_out=require_unimplemented_servers=false:$GOPATH/src proto/**/*.proto
```

Generated files will be placed in `gen/go`

---