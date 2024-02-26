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

type customerService struct {
	storage storage.IStorage
}

func NewCustomerService(storage storage.IStorage) customerService {
	return customerService{
		storage: storage,
	}
}

func (c customerService) Create(ctx context.Context, CreateCustomer models.CreateCustomer) (models.Customer, error) {

	pKey, err := c.storage.Customer().Create(ctx, CreateCustomer)
	if err != nil {
		log.Println("error in service layer while creating author ", err.Error())
		return models.Customer{}, err
	}

	customer, err := c.storage.Customer().Get(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		log.Println("error in service layer get customer by id")
	}

	return customer, nil
}

func (c customerService) Get(ctx context.Context, pkey models.PrimaryKey) (models.Customer, error) {

	customer, err := c.storage.Customer().Get(ctx, pkey)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting customer by id", err.Error())
			return models.Customer{}, err
		}
	}

	return customer, nil
}

func (c customerService) GetList(ctx context.Context, request models.GetListRequest) (models.CustomersResponse, error) {

	customers, err := c.storage.Customer().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting customers list", err.Error())
			return models.CustomersResponse{}, err
		}
	}
	return customers, nil
}

func (c customerService) Update(ctx context.Context, updateCustomer models.UpdateCustomer) (models.Customer, error) {

	id, err := c.storage.Customer().Update(ctx, updateCustomer)
	if err != nil {
		fmt.Println("error in servise layer updating customer by id", err.Error())
		return models.Customer{}, err
	}

	customer, err := c.storage.Customer().Get(ctx, models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		fmt.Println("error in service layer getting customer after update", err.Error())
		return models.Customer{}, err
	}

	return customer, nil
}

func (c customerService) Delete(ctx context.Context, id string) error {

	err := c.storage.Customer().Delete(ctx, id)

	return err
}

func (c customerService) UpdatePassword(ctx context.Context, request models.UpdateCustomerPassword) error {

	oldPassword, err := c.storage.Customer().GetPassword(ctx, request.ID)
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

	if err = c.storage.Customer().UpdatePassword(context.Background(), request); err != nil {
		fmt.Println("error in service layer while updating customer password ", err.Error())
		return err
	}

	return nil
}
