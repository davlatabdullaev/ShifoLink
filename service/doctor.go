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

type doctorService struct {
	storage storage.IStorage
}

func NewDoctorService(storage storage.IStorage) doctorService {
	return doctorService{
		storage: storage,
	}
}

func (d doctorService) Create(ctx context.Context, createDoctor models.CreateDoctor) (models.Doctor, error) {

	pKey, err := d.storage.Doctor().Create(ctx, createDoctor)
	if err != nil {
		log.Println("error in service layer while creating doctor  ", err.Error())
		return models.Doctor{}, err
	}

	doctor, err := d.storage.Doctor().Get(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		log.Println("error in service layer get doctor by id")
	}

	return doctor, nil
}

func (d doctorService) Get(ctx context.Context, pkey models.PrimaryKey) (models.Doctor, error) {

	doctor, err := d.storage.Doctor().Get(ctx, pkey)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting doctor by id", err.Error())
			return models.Doctor{}, err
		}
	}

	return doctor, nil
}

func (d doctorService) GetList(ctx context.Context, request models.GetListRequest) (models.DoctorsResponse, error) {

	doctor, err := d.storage.Doctor().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting doctor list", err.Error())
			return models.DoctorsResponse{}, err
		}
	}
	return doctor, nil
}

func (d doctorService) Update(ctx context.Context, updateDoctor models.UpdateDoctor) (models.Doctor, error) {

	id, err := d.storage.Doctor().Update(ctx, updateDoctor)
	if err != nil {
		fmt.Println("error in servise layer updating doctor type by id", err.Error())
		return models.Doctor{}, err
	}

	doctor, err := d.storage.Doctor().Get(ctx, models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		log.Println("error in service layer getting doctor after update", err.Error())
		return models.Doctor{}, err
	}

	return doctor, nil
}

func (d doctorService) Delete(ctx context.Context, id string) error {

	err := d.storage.Doctor().Delete(ctx, id)

	return err
}

func (d doctorService) UpdatePassword(ctx context.Context, request models.UpdateDoctorPassword) error {

	oldPassword, err := d.storage.Doctor().GetPassword(ctx, request.ID)
	if err != nil {
		log.Println("error in service layer getting password by id", err.Error())
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

	if err = d.storage.Doctor().UpdatePassword(context.Background(), request); err != nil {
		fmt.Println("error in service layer while updating doctor password ", err.Error())
		return err
	}

	return nil
}
