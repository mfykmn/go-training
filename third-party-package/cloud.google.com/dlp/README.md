##事前準備
1.  APIの有効化
https://console.developers.google.com/apis/api/dlp.googleapis.com/overview
2. DLP APIの権限をもたせたサービスアカウント作成
3. credential情報のJSONをこの階層にtoken.jsonという名前で設置


## 実行
export GOOGLE_APPLICATION_CREDENTIALS="token.json"
go run main.go