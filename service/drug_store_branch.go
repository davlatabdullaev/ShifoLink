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

type drugStoreBranchService struct {
	storage storage.IStorage
}

func NewDrugStoreBranchService(storage storage.IStorage) drugStoreBranchService {
	return drugStoreBranchService{
		storage: storage,
	}
}

func (d drugStoreBranchService) Create(ctx context.Context, createDrugStoreBranch models.CreateDrugStoreBranch) (models.DrugStoreBranch, error) {

	pKey, err := d.storage.DrugStoreBranch().Create(ctx, createDrugStoreBranch)
	if err != nil {
		log.Println("error in service layer while creating drug store branch  ", err.Error())
		return models.DrugStoreBranch{}, err
	}

	drugStoreBranch, err := d.storage.DrugStoreBranch().Get(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		log.Println("error in service layer get drug store branch by id")
	}

	return drugStoreBranch, nil
}

func (d drugStoreBranchService) Get(ctx context.Context, pkey models.PrimaryKey) (models.DrugStoreBranch, error) {

	drugStoreBranch, err := d.storage.DrugStoreBranch().Get(ctx, pkey)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting drug store branch by id", err.Error())
			return models.DrugStoreBranch{}, err
		}
	}

	return drugStoreBranch, nil
}

func (d drugStoreBranchService) GetList(ctx context.Context, request models.GetListRequest) (models.DrugStoreBranchsResponse, error) {

	drugStoreBranch, err := d.storage.DrugStoreBranch().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting doctor type list", err.Error())
			return models.DrugStoreBranchsResponse{}, err
		}
	}
	return drugStoreBranch, nil
}

func (d drugStoreBranchService) Update(ctx context.Context, updateDrugStoreBranch models.UpdateDrugStoreBranch) (models.DrugStoreBranch, error) {

	id, err := d.storage.DrugStoreBranch().Update(ctx, updateDrugStoreBranch)
	if err != nil {
		fmt.Println("error in servise layer updating drug store branch by id", err.Error())
		return models.DrugStoreBranch{}, err
	}

	drugStoreBranch, err := d.storage.DrugStoreBranch().Get(ctx, models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		fmt.Println("error in service layer getting doctor type after update", err.Error())
		return models.DrugStoreBranch{}, err
	}

	return drugStoreBranch, nil
}

func (d drugStoreBranchService) Delete(ctx context.Context, id string) error {

	err := d.storage.DrugStoreBranch().Delete(ctx, id)

	return err
}
