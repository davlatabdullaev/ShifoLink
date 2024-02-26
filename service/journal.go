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

type journalService struct {
	storage storage.IStorage
}

func NewJournalService(storage storage.IStorage) journalService {
	return journalService{
		storage: storage,
	}
}

func (j journalService) Create(ctx context.Context, createJournal models.CreateJournal) (models.Journal, error) {

	pKey, err := j.storage.Journal().Create(ctx, createJournal)
	if err != nil {
		log.Println("error in service layer while creating journal ", err.Error())
		return models.Journal{}, err
	}

	journal, err := j.storage.Journal().Get(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		log.Println("error in service layer get journal by id")
	}

	return journal, nil
}

func (j journalService) Get(ctx context.Context, pkey models.PrimaryKey) (models.Journal, error) {

	journal, err := j.storage.Journal().Get(ctx, pkey)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting journal by id", err.Error())
			return models.Journal{}, err
		}
	}

	return journal, nil
}

func (j journalService) GetList(ctx context.Context, request models.GetListRequest) (models.JournalsResponse, error) {

	journal, err := j.storage.Journal().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting journal  list", err.Error())
			return models.JournalsResponse{}, err
		}
	}
	return journal, nil
}

func (j journalService) Update(ctx context.Context, updateJournal models.UpdateJournal) (models.Journal, error) {

	id, err := j.storage.Journal().Update(ctx, updateJournal)
	if err != nil {
		fmt.Println("error in servise layer updating journal  by id", err.Error())
		return models.Journal{}, err
	}

	journal, err := j.storage.Journal().Get(ctx, models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		fmt.Println("error in service layer getting journal after update", err.Error())
		return models.Journal{}, err
	}

	return journal, nil
}

func (j journalService) Delete(ctx context.Context, id string) error {

	err := j.storage.Journal().Delete(ctx, id)

	return err
}
