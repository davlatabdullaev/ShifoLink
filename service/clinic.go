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

type clinicService struct {
	storage storage.IStorage
}

func NewClinicService(storage storage.IStorage) clinicService {
	return clinicService{
		storage: storage,
	}
}

func (c clinicService) Create(ctx context.Context, createClinic models.CreateClinic) (models.Clinic, error) {

	pKey, err := c.storage.Clinic().Create(ctx, createClinic)
	if err != nil {
		log.Println("error in service layer while creating clinic  ", err.Error())
		return models.Clinic{}, err
	}

	clinic, err := c.storage.Clinic().Get(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		log.Println("error in service layer get clinic by id")
	}

	return clinic, nil
}

func (c clinicService) Get(ctx context.Context, pkey models.PrimaryKey) (models.Clinic, error) {

	clinic, err := c.storage.Clinic().Get(ctx, pkey)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting clinic branch by id", err.Error())
			return models.Clinic{}, err
		}
	}

	return clinic, nil
}

func (c clinicService) GetList(ctx context.Context, request models.GetListRequest) (models.ClinicsResponse, error) {

	clinic, err := c.storage.Clinic().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting clinics list", err.Error())
			return models.ClinicsResponse{}, err
		}
	}
	return clinic, nil
}

func (c clinicService) Update(ctx context.Context, updateClinic models.UpdateClinic) (models.Clinic, error) {

	id, err := c.storage.Clinic().Update(ctx, updateClinic)
	if err != nil {
		fmt.Println("error in servise layer updating clinic by id", err.Error())
		return models.Clinic{}, err
	}

	clinic, err := c.storage.Clinic().Get(ctx, models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		fmt.Println("error in service layer getting clinic after update", err.Error())
		return models.Clinic{}, err
	}

	return clinic, nil
}

func (c clinicService) Delete(ctx context.Context, id string) error {

	err := c.storage.Clinic().Delete(ctx, id)

	return err
}
