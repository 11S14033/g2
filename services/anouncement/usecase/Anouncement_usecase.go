package usecase

import (
	"context"
	"encoding/json"
	"log"

	"github.com/11s14033/g2/services/anouncement/model"
	kafkaRepo "github.com/11s14033/g2/services/anouncement/repository/kafka"
	mongoRepo "github.com/11s14033/g2/services/anouncement/repository/mongo"
)

type anouncementUsecase struct {
	kRepo kafkaRepo.KafkaRepository
	mRepo mongoRepo.MongoRepository
}

func NewAnouncementUseCase(k kafkaRepo.KafkaRepository, m mongoRepo.MongoRepository) AnouncementUsecase {
	return &anouncementUsecase{
		kRepo: k,
		mRepo: m,
	}
}

func (aUsecase *anouncementUsecase) PublishAnouncement(ctx context.Context, an model.Anouncement, key []byte) error {
	json, err := json.Marshal(an)
	if err != nil {
		log.Printf("[Error][AnouncementUsecase][MarshalAnouncement][Cause: %v]\n", err)
	}
	err = aUsecase.kRepo.ProduceAnouncement(context.Background(), key, []byte(json))
	if err != nil {
		log.Printf("[Error][AnouncementUsecase][when call service][kafkaRepository][ProduceAnouncement][Cause: %v]\n", err)
		return err
	}

	return nil
}

func (aUsecase *anouncementUsecase) ConsumeAndInsertDB(ctx context.Context) error {

	var an model.Anouncement
	anbyte, err := aUsecase.kRepo.ConsumeAnouncement(ctx)

	if err != nil {
		log.Printf("[Error][AnouncementUsecase][ConsumeAndInsertDB][when call service][kafkaRepository][ConsumeAnouncement][Cause: %v]\n", err)
		return err
	}
	json.Unmarshal(anbyte, &an)

	err = aUsecase.mRepo.InsertAnouncement(ctx, an)
	if err != nil {
		log.Printf("[Error][AnouncementUsecase][ConsumeAndInsertDB][when call service][mongoRepository][InsertAnouncement][Cause: %v]\n", err)
		return err
	}

	return nil
}

func (aUsecase *anouncementUsecase) GetAnouncements(ctx context.Context) (ans []model.Anouncement, err error) {
	ans, err = aUsecase.mRepo.GetAnouncements(ctx)
	if err != nil {
		log.Printf("[Error][AnouncementUsecase][GetAnouncements][when call service][mongoRepository][GetAnouncements][Cause: %v]\n", err)
		return nil, err
	}

	return ans, nil
}

func (aUsecase *anouncementUsecase) GetAnouncementByType(ctx context.Context, typ string) ([]model.Anouncement, error) {
	//Next step
	return nil, nil
}
