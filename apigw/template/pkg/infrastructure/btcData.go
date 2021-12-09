package infrastructure

import (
	"template/pkg/domain"
	"template/pkg/domain/repository"
)

type btcDataPersistence struct {
}

func NewPlayDataPersistence() repository.BtcDataRepository {
	return &btcDataPersistence{}
}

//func (b btcDataPersistence) GetBTC() {
//
//}

func (b btcDataPersistence) AddBTC(req domain.BtcData) error {
	client, err := New()
	if err != nil {
		return err
	}
	err = client.Method.DynamoLogic.InsertRate(req)
	if err != nil {
		return err
	}
	return nil
}
