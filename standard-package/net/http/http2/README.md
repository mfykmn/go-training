# 自己証明書作成時に利用
$ openssl genrsa 2048 > server.key
$ openssl req -new -key server.key > server.csr
$ openssl x509 -days 3650 -req -signkey server.key < server.csr > server.crt


#参考
- https://qiita.com/koki_cheese/items/35c3fad6f1eb8458eafd
- https://deeeet.com/writing/2015/11/19/go-http2/