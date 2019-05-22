# Go
## 依存パッケージ取得
```bash
$ export GO111MODULE=on
$ go mod init
$ 
```

# AppEngine
https://cloud.google.com/appengine/docs/standard/
https://cloud.google.com/appengine/docs/quotas?hl=ja

## Config設定
```bash
$ gcloud config configurations create [CONFIG_NAME]
$ gcloud config configurations activate [CONFIG_NAME]
$ gcloud config set account [Account]
$ gcloud config set project [YOUR_PROJECT_NAME]
$ gcloud config set compute/region asia-northeast1
$ gcloud config set compute/zone asia-northeast1-a
$ gcloud config configurations list
$ gcloud auth login
```

## デプロイ
```bash
$ gcloud components install app-engine-go
$ gcloud app create --project=[YOUR_PROJECT_NAME]
$ gcloud app deploy
```

# Memo
google.golang.org/appengine packageを利用して、DatastoreやMemcacheなどを利用する場合、最初に必ず appengine.Main() を呼ぶ必要がある。