
```bash

$ cd third-party-package/jwt-go/
$ docker-compose up -d
$ export SIGNINGKEY=aaaaa
$ realize start --server

$  curl -H "Authorization:Bearer {JWTトークン}" localhost:8080
```
