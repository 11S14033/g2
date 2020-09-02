package mongo

import (
	"context"
	"log"

	"github.com/11s14033/g2/services/anouncement/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(col *mongo.Collection) MongoRepository {
	return &mongoRepository{
		collection: col,
	}
}

func (m *mongoRepository) GetAnouncements(ctx context.Context) ([]model.Anouncement, error) {
	an := model.Anouncement{}
	ans := []model.Anouncement{}
	cursor, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("[Error][MongoRepository][GetAnouncements][cause: %v]\n", err)
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		err := cursor.Decode(&an)
		if err != nil {
			log.Printf("[Error][MongoRepository][GetAnouncements][Decode data][cause: %v]\n", err)
			return nil, err
		}
		ans = append(ans, an)

	}

	return ans, nil
}

func (m *mongoRepository) GetAnouncementByType(ctx context.Context) ([]model.Anouncement, error) {
	//Next step
	return nil, nil
}

func (m *mongoRepository) InsertAnouncement(ctx context.Context, an model.Anouncement) error {
	cur, err := m.collection.InsertOne(ctx, an)
	if err != nil {
		log.Printf("[Error][MongoRepository][InsertAnouncement][cause: %v]\n", err)
		return err
	}

	log.Printf("[MongoRepository][InsertAnouncement][Insert ID : %v]\n", cur.InsertedID)
	return nil
}

func (m *mongoRepository) InsertBatchAnouncement(ctx context.Context, ans []model.Anouncement) error {
	//Next step
	return nil
}
