# pancake

```bash
$ brew install protobuf
$ go get -u github.com/golang/protobuf/protoc-gen-go
$ protoc \
  -Iproto \
  --go_out=plugins=grpc:. \
  proto/*.proto
```