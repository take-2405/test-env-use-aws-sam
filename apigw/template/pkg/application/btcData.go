package application

import (
	"template/pkg/domain"
	"template/pkg/domain/repository"
)

type BtcUseCase interface {
	InsertRateUseCase(req domain.BtcData) error
}

type btcUseCase struct {
	btcRepository repository.BtcDataRepository
}

func NewBtcUseCase(b repository.BtcDataRepository) BtcUseCase {
	return &btcUseCase{btcRepository: b}
}

func (bu btcUseCase) InsertRateUseCase(req domain.BtcData) error {
	err := bu.btcRepository.AddBTC(req)
	if err != nil {
		return err
	}
	return nil
}
