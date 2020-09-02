package mongo

import (
	"context"

	"github.com/11s14033/g2/services/anouncement/model"
)

type MongoRepository interface {
	GetAnouncements(ctx context.Context) ([]model.Anouncement, error)
	GetAnouncementByType(ctx context.Context) ([]model.Anouncement, error)
	InsertAnouncement(ctx context.Context, an model.Anouncement) error
	InsertBatchAnouncement(ctx context.Context, ans []model.Anouncement) error
}
