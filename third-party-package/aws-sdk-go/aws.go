package mock

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	accessKeyID     = ""
	secretAccessKey = ""
	region          = "us-west-2"
)

func InitAWSSession() *session.Session {
	// クレデンシャルの作成
	cred := credentials.NewStaticCredentials(accessKeyID, secretAccessKey, "") // 最後の引数は[セッショントークン]

	// クレデンシャルとリージョンをセットしたコンフィグの作成
	conf := &aws.Config{
		Credentials: cred,
		Region:      &region,
	}

	// セッションの作成
	sess, err := session.NewSession(conf)
	if err != nil {
		panic(err)
	}
	return sess
}
