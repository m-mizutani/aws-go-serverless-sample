package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/sirupsen/logrus"
)

type LambdaArguments struct {
	Region string
	Event  interface{}
}

var (
	loggerBase = logrus.New()
	logger     *logrus.Entry
)

func MainHandler(args LambdaArguments) error {
	// ここにメインの処理を書く
	return nil
}

func handleRequest(ctx context.Context, event interface{}) error {
	// デプロイ後の環境で必要になる処理のセットアップ等をする
	loggerBase.SetFormatter(&logrus.JSONFormatter{})
	loggerBase.SetLevel(logrus.InfoLevel)
	lc, _ := lambdacontext.FromContext(ctx)
	logger = loggerBase.WithField("request_id", lc.AwsRequestID)

	args := LambdaArguments{
		Event:  event,
		Region: os.Getenv("AWS_REGION"),
	}

	if err := MainHandler(args); err != nil {
		// なんかエラー処理をする
		return err
	}

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
