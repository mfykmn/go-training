# Go
## 依存パッケージ取得
```bash
$ export GO111MODULE=on
$ go mod init
$ 
```

# AppEngine
## Config設定
```bash
$ gcloud config configurations create {name}
$ gcloud config configurations activate  {config_name}
$ gcloud config set account {account}
$ gcloud config set project {project_id}
$ gcloud config set compute/region asia-northeast1
$ gcloud config set compute/zone asia-northeast1-a
$ gcloud config configurations list
$ gcloud auth login
```

## デプロイ
```bash
$ gcloud app create
$ gcloud app deploy
```

# Memo
google.golang.org/appengine packageを利用して、DatastoreやMemcacheなどを利用する場合、最初に必ず appengine.Main() を呼ぶ必要がある。