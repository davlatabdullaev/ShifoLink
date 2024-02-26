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

type clinicAdminService struct {
	storage storage.IStorage
}

func NewClinicAdminService(storage storage.IStorage) clinicAdminService {
	return clinicAdminService{
		storage: storage,
	}
}

func (c clinicAdminService) Create(ctx context.Context, createClinicAdmin models.CreateClinicAdmin) (models.ClinicAdmin, error) {

	pKey, err := c.storage.ClinicAdmin().Create(ctx, createClinicAdmin)
	if err != nil {
		log.Println("error in service layer while creating clinic admin ", err.Error())
		return models.ClinicAdmin{}, err
	}

	clinicAdmin, err := c.storage.ClinicAdmin().Get(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		log.Println("error in service layer get clinic admin by id")
	}

	return clinicAdmin, nil
}

func (c clinicAdminService) Get(ctx context.Context, pkey models.PrimaryKey) (models.ClinicAdmin, error) {

	clinicAdmin, err := c.storage.ClinicAdmin().Get(ctx, pkey)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting clinic admin by id", err.Error())
			return models.ClinicAdmin{}, err
		}
	}

	return clinicAdmin, nil
}

func (c clinicAdminService) GetList(ctx context.Context, request models.GetListRequest) (models.ClinicAdminsResponse, error) {

	clinicAdmin, err := c.storage.ClinicAdmin().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting clinic admin list", err.Error())
			return models.ClinicAdminsResponse{}, err
		}
	}
	return clinicAdmin, nil
}

func (c clinicAdminService) Update(ctx context.Context, updateClinicAdmin models.UpdateClinicAdmin) (models.ClinicAdmin, error) {

	id, err := c.storage.ClinicAdmin().Update(ctx, updateClinicAdmin)
	if err != nil {
		fmt.Println("error in servise layer updating clinic admin by id", err.Error())
		return models.ClinicAdmin{}, err
	}

	clinicAdmin, err := c.storage.ClinicAdmin().Get(ctx, models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		fmt.Println("error in service layer getting clinic admin after update", err.Error())
		return models.ClinicAdmin{}, err
	}

	return clinicAdmin, nil
}

func (c clinicAdminService) Delete(ctx context.Context, id string) error {

	err := c.storage.ClinicAdmin().Delete(ctx, id)

	return err
}
