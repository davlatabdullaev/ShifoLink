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

type drugStoreService struct {
	storage storage.IStorage
}

func NewDrugStoreService(storage storage.IStorage) drugStoreService {
	return drugStoreService{
		storage: storage,
	}
}

func (d drugStoreService) Create(ctx context.Context, createDrugStore models.CreateDrugStore) (models.DrugStore, error) {

	pKey, err := d.storage.DrugStore().Create(ctx, createDrugStore)
	if err != nil {
		log.Println("error in service layer while creating drug store   ", err.Error())
		return models.DrugStore{}, err
	}

	drugStore, err := d.storage.DrugStore().Get(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		log.Println("error in service layer get drug store by id")
	}

	return drugStore, nil
}

func (d drugStoreService) Get(ctx context.Context, pkey models.PrimaryKey) (models.DrugStore, error) {

	drugStore, err := d.storage.DrugStore().Get(ctx, pkey)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting drug store by id", err.Error())
			return models.DrugStore{}, err
		}
	}

	return drugStore, nil
}

func (d drugStoreService) GetList(ctx context.Context, request models.GetListRequest) (models.DrugStoresResponse, error) {

	drugStore, err := d.storage.DrugStore().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting drug store list", err.Error())
			return models.DrugStoresResponse{}, err
		}
	}
	return drugStore, nil
}

func (d drugStoreService) Update(ctx context.Context, updateDrugStore models.UpdateDrugStore) (models.DrugStore, error) {

	id, err := d.storage.DrugStore().Update(ctx, updateDrugStore)
	if err != nil {
		fmt.Println("error in servise layer updating drug store by id", err.Error())
		return models.DrugStore{}, err
	}

	drugStore, err := d.storage.DrugStore().Get(ctx, models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		fmt.Println("error in service layer getting drug store after update", err.Error())
		return models.DrugStore{}, err
	}

	return drugStore, nil
}

func (d drugStoreService) Delete(ctx context.Context, id string) error {

	err := d.storage.DrugStore().Delete(ctx, id)

	return err
}
