package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"template/pkg/domain"
)

type DynamoDB struct {
	Dynamo *dynamodb.DynamoDB
	Method Methods
}

type Methods struct {
	DynamoLogic methods
}

type methods interface {
	//GetRate()
	InsertRate(req domain.BtcData) error
}

func New() (*DynamoDB, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://dynamoTest:8000"),
		Region:   aws.String("ap-northeast-1"),
	}))

	//DB接続
	svc := dynamodb.New(sess)

	// init methods
	methods := newDynamoClient(svc)

	return &DynamoDB{
		Dynamo: svc,
		Method: Methods{methods},
	}, nil
}
