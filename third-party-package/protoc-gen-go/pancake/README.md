# pancake

```bash
$ brew install protobuf
$ go get -u github.com/golang/protobuf/protoc-gen-go
$ protoc \
  -Iproto \
  --go_out=plugins=grpc:. \
  proto/*.proto
```


```bash
$ brew tap grpc/grpc
$ brew install grpc
```

```bash
$ go run server.go baker_handler.go
$ grpc_cli ls localhost:50051 pancake.baker.PancakeBakerService -l
 filename: pancake.proto
 package: pancake.baker;
 service PancakeBakerService {
   rpc Bake(pancake.baker.BakeRequest) returns (pancake.baker.BakeResponse) {}
   rpc Report(pancake.baker.ReportRequest) returns (pancake.baker.ReportResponse) {}
 }

$ grpc_cli call localhost:50051 pancake.baker.PancakeBakerService.Bake 'menu: 1'
  connecting to localhost:50051
  pancake {
    chef_name: "gami"
    menu: CLASSIC
    technical_score: 0.860443115
    create_time {
      seconds: 1569399934
      nanos: 532942000
    }
  }
  
  Rpc succeeded with OK status
```