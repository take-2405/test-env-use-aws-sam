package domain

type BtcData struct {
	Rate      int `dynamodbav:"rate"`
	Timestamp int `dynamodbav:"timestamp"`
}
