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

type superAdminService struct {
	storage storage.IStorage
}

func NewSuperAdminService(storage storage.IStorage) superAdminService {
	return superAdminService{
		storage: storage,
	}
}

func (s superAdminService) Create(ctx context.Context, createSuperAdmin models.CreateSuperAdmin) (models.SuperAdmin, error) {

	pKey, err := s.storage.SuperAdmin().Create(ctx, createSuperAdmin)
	if err != nil {
		log.Println("error in service layer while creating superAdmin ", err.Error())
		return models.SuperAdmin{}, err
	}

	superAdmin, err := s.storage.SuperAdmin().Get(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		log.Println("error in service layer get superAdmin by id")
	}

	return superAdmin, nil
}

func (s superAdminService) Get(ctx context.Context, pkey models.PrimaryKey) (models.SuperAdmin, error) {

	superAdmin, err := s.storage.SuperAdmin().Get(ctx, pkey)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting superAdmin by id", err.Error())
			return models.SuperAdmin{}, err
		}
	}

	return superAdmin, nil
}

func (s superAdminService) GetList(ctx context.Context, request models.GetListRequest) (models.SuperAdminsResponse, error) {

	superAdmin, err := s.storage.SuperAdmin().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting superAdmin list", err.Error())
			return models.SuperAdminsResponse{}, err
		}
	}
	return superAdmin, nil
}

func (s superAdminService) Update(ctx context.Context, updateSuperAdmin models.UpdateSuperAdmin) (models.SuperAdmin, error) {

	id, err := s.storage.SuperAdmin().Update(ctx, updateSuperAdmin)
	if err != nil {
		fmt.Println("error in servise layer updating superAdmin by id", err.Error())
		return models.SuperAdmin{}, err
	}

	superAdmin, err := s.storage.SuperAdmin().Get(ctx, models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		fmt.Println("error in service layer getting superAdmin after update", err.Error())
		return models.SuperAdmin{}, err
	}

	return superAdmin, nil
}

func (s superAdminService) Delete(ctx context.Context, id string) error {

	err := s.storage.SuperAdmin().Delete(ctx, id)

	return err
}

func (s superAdminService) UpdatePassword(ctx context.Context, request models.UpdateSuperAdminPassword) error {

	oldPassword, err := s.storage.SuperAdmin().GetPassword(ctx, request.ID)
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

	if err = s.storage.SuperAdmin().UpdatePassword(context.Background(), request); err != nil {
		fmt.Println("error in service layer while updating pharmacist password ", err.Error())
		return err
	}

	return nil
}
