# Protocol Buffers
GolangでProtocol Buffersを試してみる

## 事前準備

```bash
$ brew install protobuf
$ go get -u github.com/golang/protobuf/protoc-gen-go
$ cd cd third-party-package/protobuf/
$ protoc  --go_out=./ ./addressbook.proto
```

## 参考リンク
- [Googleのチュートリアル](https://developers.google.com/protocol-buffers/docs/gotutorial)
- [今さらProtocol Buffersと、手に馴染む道具の話](https://qiita.com/yugui/items/160737021d25d761b353#fn1)
- [ProtocolBuffersについて調べてみた](https://qiita.com/aiueo4u/items/54dc5dd8c4772253634c)
- [gRPC-Web is going GA](https://www.cncf.io/blog/2018/10/24/grpc-web-is-going-ga/)