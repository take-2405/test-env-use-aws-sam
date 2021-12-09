package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"template/pkg/application"
	"template/pkg/domain/repository"
	"template/pkg/infrastructure"
	"template/pkg/presentation"
)

func main() {
	btcPersistence := infrastructure.NewPlayDataPersistence()
	btcRepository := repository.BtcDataRepository(btcPersistence)
	btcUseCase := application.NewBtcUseCase(btcRepository)
	handler := presentation.NewBTCHandler(btcUseCase)
	lambda.Start(handler.HandleBtcRanking)
}
