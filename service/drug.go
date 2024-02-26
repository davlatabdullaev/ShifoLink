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

type drugService struct {
	storage storage.IStorage
}

func NewDrugService(storage storage.IStorage) drugService {
	return drugService{
		storage: storage,
	}
}

func (d drugService) Create(ctx context.Context, createDrug models.CreateDrug) (models.Drug, error) {

	pKey, err := d.storage.Drug().Create(ctx, createDrug)
	if err != nil {
		log.Println("error in service layer while creating drug ", err.Error())
		return models.Drug{}, err
	}

	drug, err := d.storage.Drug().Get(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		log.Println("error in service layer get drug by id")
	}

	return drug, nil
}

func (d drugService) Get(ctx context.Context, pkey models.PrimaryKey) (models.Drug, error) {

	drug, err := d.storage.Drug().Get(ctx, pkey)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting drug by id", err.Error())
			return models.Drug{}, err
		}
	}

	return drug, nil
}

func (d drugService) GetList(ctx context.Context, request models.GetListRequest) (models.DrugsResponse, error) {

	drug, err := d.storage.Drug().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting drug  list", err.Error())
			return models.DrugsResponse{}, err
		}
	}
	return drug, nil
}

func (d drugService) Update(ctx context.Context, updateDrug models.UpdateDrug) (models.Drug, error) {

	id, err := d.storage.Drug().Update(ctx, updateDrug)
	if err != nil {
		fmt.Println("error in servise layer updating drug  by id", err.Error())
		return models.Drug{}, err
	}

	drug, err := d.storage.Drug().Get(ctx, models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		fmt.Println("error in service layer getting drug after update", err.Error())
		return models.Drug{}, err
	}

	return drug, nil
}

func (d drugService) Delete(ctx context.Context, id string) error {

	err := d.storage.Drug().Delete(ctx, id)

	return err
}
