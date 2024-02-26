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

type doctorTypeService struct {
	storage storage.IStorage
}

func NewDoctorTypeService(storage storage.IStorage) doctorTypeService {
	return doctorTypeService{
		storage: storage,
	}
}

func (d doctorTypeService) Create(ctx context.Context, createDoctorType models.CreateDoctorType) (models.DoctorType, error) {

	pKey, err := d.storage.DoctorType().Create(ctx, createDoctorType)
	if err != nil {
		log.Println("error in service layer while creating clinic  ", err.Error())
		return models.DoctorType{}, err
	}

	doctorType, err := d.storage.DoctorType().Get(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		log.Println("error in service layer get doctor type by id")
	}

	return doctorType, nil
}

func (d doctorTypeService) Get(ctx context.Context, pkey models.PrimaryKey) (models.DoctorType, error) {

	doctorType, err := d.storage.DoctorType().Get(ctx, pkey)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting doctor type by id", err.Error())
			return models.DoctorType{}, err
		}
	}

	return doctorType, nil
}

func (d doctorTypeService) GetList(ctx context.Context, request models.GetListRequest) (models.DoctorTypesResponse, error) {

	doctorType, err := d.storage.DoctorType().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting doctor type list", err.Error())
			return models.DoctorTypesResponse{}, err
		}
	}
	return doctorType, nil
}

func (d doctorTypeService) Update(ctx context.Context, updateDoctorType models.UpdateDoctorType) (models.DoctorType, error) {

	id, err := d.storage.DoctorType().Update(ctx, updateDoctorType)
	if err != nil {
		fmt.Println("error in servise layer updating doctor type by id", err.Error())
		return models.DoctorType{}, err
	}

	doctorType, err := d.storage.DoctorType().Get(ctx, models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		fmt.Println("error in service layer getting doctor type after update", err.Error())
		return models.DoctorType{}, err
	}

	return doctorType, nil
}

func (d doctorTypeService) Delete(ctx context.Context, id string) error {

	err := d.storage.DoctorType().Delete(ctx, id)

	return err
}
