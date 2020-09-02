package usecase

import (
	"context"

	"github.com/11s14033/g2/services/anouncement/model"
)

type AnouncementUsecase interface {
	PublishAnouncement(ctx context.Context, an model.Anouncement, key []byte) error
	ConsumeAndInsertDB(ctx context.Context) error
	GetAnouncements(ctx context.Context) ([]model.Anouncement, error)
	GetAnouncementByType(ctx context.Context, typ string) ([]model.Anouncement, error)
}
