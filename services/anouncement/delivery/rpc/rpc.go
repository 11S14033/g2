package rpc

import (
	aUsecase "github.com/11s14033/g2/services/anouncement/usecase"
)

type AnouncementRPC struct {
	aUS aUsecase.AnouncementUsecase
}

func NewAnouncementRPC(a aUsecase.AnouncementUsecase) AnouncementRPC {
	return AnouncementRPC{
		aUS: a,
	}
}
