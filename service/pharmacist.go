package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"shifolink/api/models"
	"shifolink/pkg/check"
	"shifolink/storage"

	"github.com/jackc/pgx/v5"
)

type pharmacistService struct {
	storage storage.IStorage
}

func NewPharmacistService(storage storage.IStorage) pharmacistService {
	return pharmacistService{
		storage: storage,
	}
}

func (p pharmacistService) Create(ctx context.Context, createPharmacist models.CreatePharmacist) (models.Pharmacist, error) {

	pKey, err := p.storage.Pharmacist().Create(ctx, createPharmacist)
	if err != nil {
		log.Println("error in service layer while creating pharmacist ", err.Error())
		return models.Pharmacist{}, err
	}

	pharmacist, err := p.storage.Pharmacist().Get(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		log.Println("error in service layer get pharmacist by id")
	}

	return pharmacist, nil
}

func (p pharmacistService) Get(ctx context.Context, pkey models.PrimaryKey) (models.Pharmacist, error) {

	pharmacist, err := p.storage.Pharmacist().Get(ctx, pkey)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting pharmacist by id", err.Error())
			return models.Pharmacist{}, err
		}
	}

	return pharmacist, nil
}

func (p pharmacistService) GetList(ctx context.Context, request models.GetListRequest) (models.PharmacistsResponse, error) {

	customers, err := p.storage.Pharmacist().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting customers list", err.Error())
			return models.PharmacistsResponse{}, err
		}
	}
	return customers, nil
}

func (p pharmacistService) Update(ctx context.Context, updatePharmacist models.UpdatePharmacist) (models.Pharmacist, error) {

	id, err := p.storage.Pharmacist().Update(ctx, updatePharmacist)
	if err != nil {
		fmt.Println("error in servise layer updating pharmacist by id", err.Error())
		return models.Pharmacist{}, err
	}

	pharmacist, err := p.storage.Pharmacist().Get(ctx, models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		fmt.Println("error in service layer getting pharmacist after update", err.Error())
		return models.Pharmacist{}, err
	}

	return pharmacist, nil
}

func (p pharmacistService) Delete(ctx context.Context, id string) error {

	err := p.storage.Pharmacist().Delete(ctx, id)

	return err
}

func (p pharmacistService) UpdatePassword(ctx context.Context, request models.UpdatePharmacistPassword) error {

	oldPassword, err := p.storage.Pharmacist().GetPassword(ctx, request.ID)
	if err != nil {
		fmt.Println("error in service layer getting password by id", err.Error())
		return err
	}

	if oldPassword != request.OldPassword {
		fmt.Println("error in service layer old password is not correct")
		return errors.New("old password did not match")
	}

	if err = check.ValidatePassword(request.NewPassword); err != nil {
		fmt.Println("error in service layer new password validation failed", err.Error())
		return err
	}

	if err = p.storage.Pharmacist().UpdatePassword(context.Background(), request); err != nil {
		fmt.Println("error in service layer while updating pharmacist password ", err.Error())
		return err
	}

	return nil
}
