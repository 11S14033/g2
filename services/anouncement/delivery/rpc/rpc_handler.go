package rpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/11s14033/g2/services/anouncement/model"

	"github.com/11s14033/g2/commons/pb"
)

func (arpc AnouncementRPC) PublishAnouncement(ctx context.Context, req *pb.PublishAnouncementReq) (*pb.PublishAnouncementRes, error) {
	data := req.GetAnountcement()

	a := model.Anouncement{
		Title:     data.GetTitle(),
		Message:   data.GetMessage(),
		Author:    data.GetAuthor(),
		Type:      data.GetType(),
		CreatedAt: time.Now(),
	}
	err := arpc.aUS.PublishAnouncement(context.Background(), a, []byte("test"))
	if err != nil {
		log.Println("[Error][rpc][when call service][anouncementuseCase][PublishAnouncement][cause: %v]\n", err)
		return nil, err
	}

	return &pb.PublishAnouncementRes{
		Status: "Succes send to kafka",
	}, nil
}

func (arpc AnouncementRPC) GetAnouncements(req *pb.GetAnouncementsReq, stream pb.AnouncementServices_GetAnouncementsServer) error {
	for {

		ans, err := arpc.aUS.GetAnouncements(context.Background())

		if err != nil {
			log.Println("[Error][rpc][GetAnouncements][when call service][anouncementUseCase][GetAnouncements][cause: %v]\n", err)
			return err
		}
		for _, an := range ans {
			stream.Send(&pb.GetAnouncementsRes{
				Anountcement: &pb.Anountcement{
					Author:  an.Author,
					Message: an.Message,
					Title:   an.Title,
					Type:    an.Type,
				},
			})

		}
	}

	return nil
}

func (arpc AnouncementRPC) ConsumeAndSave(req *pb.ConsumeAndSaveReq, stream pb.AnouncementServices_ConsumeAndSaveServer) error {
	fmt.Println("start consume kafka")
	for {
		err := arpc.aUS.ConsumeAndInsertDB(context.Background())
		if err != nil {
			log.Println("[Error][rpc][ConsumeAndSave][when call service][anouncement][ConsumeAndInsertDB][cause: %v]\n", err)
			return err
		}

		stream.Send(&pb.ConsumeAndSaveRes{
			Status: "Success consume and save DB",
		})
	}

	return nil
}
