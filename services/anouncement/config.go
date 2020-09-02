package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"

	pb "github.com/11s14033/g2/commons/pb"
	rpc "github.com/11s14033/g2/services/anouncement/delivery/rpc"
	kafkaRepo "github.com/11s14033/g2/services/anouncement/repository/kafka"
	mongoRepo "github.com/11s14033/g2/services/anouncement/repository/mongo"
	usecase "github.com/11s14033/g2/services/anouncement/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type service struct {
	collection *mongo.Collection
	producer   *kafka.Writer
	consumer   *kafka.Reader
}

func initKafkaProducer() *kafka.Writer {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: "",
	}
	config := kafka.WriterConfig{
		Brokers:          []string{"localhost:9092"},
		Topic:            "diantest",
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	w := kafka.NewWriter(config)

	return w

}

func initKafkaConsume() *kafka.Reader {
	config := kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "diantest",
		GroupID: "Anouncement",
	}

	r := kafka.NewReader(config)

	return r
}

func initMongoDB() *mongo.Collection {
	fmt.Println("Connecting to mongo DB")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	collection := client.Database("mydb").Collection("anouncement")

	return collection
}

func StartGRPCServer(rpcServer pb.AnouncementServicesServer) error {
	//starting grpc server
	fmt.Println("Starting grpc server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error starting grpc cause: %v", err)
	}

	opts := []grpc.ServerOption{}

	svr := grpc.NewServer(opts...)

	pb.RegisterAnouncementServicesServer(svr, rpcServer)

	reflection.Register(svr)

	return svr.Serve(lis)
}

func (s *service) StartService() error {
	var err error
	s.producer = initKafkaProducer()
	s.consumer = initKafkaConsume()
	s.collection = initMongoDB()
	kRepo := kafkaRepo.NewKafkaRepository(s.producer, s.consumer)
	mRepo := mongoRepo.NewMongoRepository(s.collection)
	aUsecase := usecase.NewAnouncementUseCase(kRepo, mRepo)
	rpcServer := rpc.NewAnouncementRPC(aUsecase)

	//start consumer and insert to DB

	err = aUsecase.ConsumeAndInsertDB(context.Background())
	if err != nil {
		return err
	}

	err = StartGRPCServer(rpcServer)
	if err != nil {
		log.Fatalf("[Main][Error when starting gRPC server][cause:%v] ", err)
		return err
	}

	return nil

}
