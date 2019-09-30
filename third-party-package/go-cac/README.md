# go-cache

# 利用パッケージ
https://github.com/patrickmn/go-cac


curl
=== 

# IDマッピング取得
```
curl -X GET http://localhost:60001/cache/id-mapping/ID00001
```

# IDマッピングキャッシュ
```
curl -X POST http://localhost:60001/cache/id-mapping -d '{"mid":"ID00001","oid":"aaaaaaaa"}' -H "Content-type: application/json"
```
キャッシュした情報は10秒間キャッシュされる
