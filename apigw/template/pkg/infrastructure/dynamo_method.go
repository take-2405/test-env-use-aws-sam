package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"template/pkg/domain"
)

type dynamoMethods struct {
	Client *dynamodb.DynamoDB
}

func newDynamoClient(client *dynamodb.DynamoDB) methods {
	return &dynamoMethods{Client: client}
}

//func (d *dynamoMethods) GetRate() {
//
//}

func (d *dynamoMethods) InsertRate(req domain.BtcData) error {
	insert, err := dynamodbattribute.MarshalMap(req)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	putParams := &dynamodb.PutItemInput{
		TableName: aws.String("btc-rate"),
		Item:      insert,
	}

	//Execute.
	_, err = d.Client.PutItem(putParams)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
