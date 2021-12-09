package presentation

import (
	"github.com/aws/aws-lambda-go/events"
	"log"
	"template/pkg/application"
	"template/pkg/domain"
)

type BTCHandler interface {
	HandleBtcRanking(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

type btcHandler struct {
	btcUseCase application.BtcUseCase
}

func NewBTCHandler(bu application.BtcUseCase) BTCHandler {
	return &btcHandler{
		btcUseCase: bu,
	}
}

func (bh btcHandler) HandleBtcRanking(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req := domain.BtcData{Rate: 5467470, Timestamp: 1638997225}
	log.Println("おっけい")
	err := bh.btcUseCase.InsertRateUseCase(req)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
				"Content-Type":                 "application/json",
			},
			Body:       err.Error(),
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
			"Content-Type":                 "application/json",
		},
		Body:       "ok",
		StatusCode: 200,
	}, nil

}
