package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type kafkaRepository struct {
	producer *kafka.Writer // to produce
	consumer *kafka.Reader // to consume
}

func NewKafkaRepository(p *kafka.Writer, c *kafka.Reader) KafkaRepository {
	return &kafkaRepository{
		producer: p,
		consumer: c,
	}
}

func (k *kafkaRepository) ProduceAnouncement(ctx context.Context, key, value []byte) error {

	msg := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}

	err := k.producer.WriteMessages(ctx, msg)
	if err != nil {
		log.Printf("[Error][KafkaRepository][ProduceAnouncement][cause: %v]\n", err)
		return err

	}

	return nil

}
func (k *kafkaRepository) ConsumeAnouncement(ctx context.Context) (Anouncement []byte, err error) {
	msg, err := k.consumer.ReadMessage(ctx)
	if err != nil {
		log.Printf("[Error][MongoRepository][ConsumeAnouncement][cause: %v]\n", err)
		return nil, err
	}

	Anouncement = msg.Value

	return Anouncement, nil
}
func (k *kafkaRepository) ConsumeBatchAnouncement(ctx context.Context) (Anouncement []byte, err error) {
	//Next step
	return nil, nil
}
