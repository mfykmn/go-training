# 使い方
```
$ make dstart
$ make dlog # app1を選択

# 別窓で 
$ make dlog # app2を選択

# subscription
$ curl http://localhost:8081/sub
$ curl http://localhost:8082/sub

# publish
$ curl http://localhost:8082/pub # 8081にも届いている

```

#PUB/SUB
- https://redis.io/topics/pubsub
- https://qiita.com/jkr_2255/items/6ae9e5e2244052e41369
- http://need4answer.blogspot.com/2015/05/golangredis.html
- https://blog.dakatsuka.jp/2011/06/19/nodejs-redis-pubsub.html
- http://redis.shibu.jp/commandreference/pubsub.html#command-PUBLISH
- https://godoc.org/github.com/gomodule/redigo/redis#PubSubConn.PSubscribe
