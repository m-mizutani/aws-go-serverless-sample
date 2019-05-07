package main_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	main "github.com/m-mizutani/aws-go-serverless-sample"
)

func TestFuncA(t *testing.T) {
	args := main.LambdaArguments{
		Event:  nil,
		Region: os.Getenv("AWS_REGION"),
	}

	err := main.MainHandler(args)
	assert.NoError(t, err)
}
