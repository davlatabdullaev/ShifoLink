package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"shifolink/api/models"
	"shifolink/storage"

	"github.com/jackc/pgx/v5"
)

type queueService struct {
	storage storage.IStorage
}

func NewQueueService(storage storage.IStorage) queueService {
	return queueService{
		storage: storage,
	}
}

func (q queueService) Create(ctx context.Context, createQueue models.CreateQueue) (models.Queue, error) {

	pKey, err := q.storage.Queue().Create(ctx, createQueue)
	if err != nil {
		log.Println("error in service layer while creating queue ", err.Error())
		return models.Queue{}, err
	}

	queue, err := q.storage.Queue().Get(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		log.Println("error in service layer get queue by id")
	}

	return queue, nil
}

func (q queueService) Get(ctx context.Context, pkey models.PrimaryKey) (models.Queue, error) {

	queue, err := q.storage.Queue().Get(ctx, pkey)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting queue by id", err.Error())
			return models.Queue{}, err
		}
	}

	return queue, nil
}

func (q queueService) GetList(ctx context.Context, request models.GetListRequest) (models.QueuesResponse, error) {

	queue, err := q.storage.Queue().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting queue  list", err.Error())
			return models.QueuesResponse{}, err
		}
	}
	return queue, nil
}

func (q queueService) Update(ctx context.Context, updateQueue models.UpdateQueue) (models.Queue, error) {

	id, err := q.storage.Queue().Update(ctx, updateQueue)
	if err != nil {
		fmt.Println("error in servise layer updating queue  by id", err.Error())
		return models.Queue{}, err
	}

	queue, err := q.storage.Queue().Get(ctx, models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		fmt.Println("error in service layer getting queue after update", err.Error())
		return models.Queue{}, err
	}

	return queue, nil
}

func (q queueService) Delete(ctx context.Context, id string) error {

	err := q.storage.Queue().Delete(ctx, id)

	return err
}
