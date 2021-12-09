package repository

import "template/pkg/domain"

type BtcDataRepository interface {
	//GetBTC()
	AddBTC(req domain.BtcData) error
}
