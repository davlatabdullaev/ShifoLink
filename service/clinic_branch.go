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

type clinicBranchService struct {
	storage storage.IStorage
}

func NewClinicBranchService(storage storage.IStorage) clinicBranchService {
	return clinicBranchService{
		storage: storage,
	}
}

func (c clinicBranchService) Create(ctx context.Context, createClinicBranch models.CreateClinicBranch) (models.ClinicBranch, error) {

	pKey, err := c.storage.ClinicBranch().Create(ctx, createClinicBranch)
	if err != nil {
		log.Println("error in service layer while creating clinic branch ", err.Error())
		return models.ClinicBranch{}, err
	}

	clinicBranch, err := c.storage.ClinicBranch().Get(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		log.Println("error in service layer get clinic branch by id")
	}

	return clinicBranch, nil
}

func (c clinicBranchService) Get(ctx context.Context, pkey models.PrimaryKey) (models.ClinicBranch, error) {

	clinicBranch, err := c.storage.ClinicBranch().Get(ctx, pkey)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting clinic branch by id", err.Error())
			return models.ClinicBranch{}, err
		}
	}

	return clinicBranch, nil
}

func (c clinicBranchService) GetList(ctx context.Context, request models.GetListRequest) (models.ClinicBranchsResponse, error) {

	clinicBranch, err := c.storage.ClinicBranch().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting clinic branch list", err.Error())
			return models.ClinicBranchsResponse{}, err
		}
	}
	return clinicBranch, nil
}

func (c clinicBranchService) Update(ctx context.Context, updateClinicBranch models.UpdateClinicBranch) (models.ClinicBranch, error) {

	id, err := c.storage.ClinicBranch().Update(ctx, updateClinicBranch)
	if err != nil {
		fmt.Println("error in servise layer updating clinic branch by id", err.Error())
		return models.ClinicBranch{}, err
	}

	clinicBranch, err := c.storage.ClinicBranch().Get(ctx, models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		fmt.Println("error in service layer getting clinic branch after update", err.Error())
		return models.ClinicBranch{}, err
	}

	return clinicBranch, nil
}

func (c clinicBranchService) Delete(ctx context.Context, id string) error {

	err := c.storage.ClinicBranch().Delete(ctx, id)

	return err
}
