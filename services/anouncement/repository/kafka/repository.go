package kafka

import "context"

type KafkaRepository interface {
	ProduceAnouncement(ctx context.Context, key, value []byte) error
	ConsumeAnouncement(ctx context.Context) (Anouncement []byte, err error)
	ConsumeBatchAnouncement(ctx context.Context) (Anouncement []byte, err error)
}
