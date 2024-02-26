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

type orderDrugService struct {
	storage storage.IStorage
}

func NewOrderDrugService(storage storage.IStorage) orderDrugService {
	return orderDrugService{
		storage: storage,
	}
}

func (o orderDrugService) Create(ctx context.Context, createOrderDrug models.CreateOrderDrug) (models.OrderDrug, error) {

	pKey, err := o.storage.OrderDrug().Create(ctx, createOrderDrug)
	if err != nil {
		log.Println("error in service layer while creating order drug ", err.Error())
		return models.OrderDrug{}, err
	}

	orderDrug, err := o.storage.OrderDrug().Get(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		log.Println("error in service layer get order drug by id")
	}

	return orderDrug, nil
}

func (o orderDrugService) Get(ctx context.Context, pkey models.PrimaryKey) (models.OrderDrug, error) {

	orderDrug, err := o.storage.OrderDrug().Get(ctx, pkey)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting order drug by id", err.Error())
			return models.OrderDrug{}, err
		}
	}

	return orderDrug, nil
}

func (o orderDrugService) GetList(ctx context.Context, request models.GetListRequest) (models.OrderDrugsResponse, error) {

	orderDrug, err := o.storage.OrderDrug().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting order drug  list", err.Error())
			return models.OrderDrugsResponse{}, err
		}
	}
	return orderDrug, nil
}

func (o orderDrugService) Update(ctx context.Context, updateOrderDrug models.UpdateOrderDrug) (models.OrderDrug, error) {

	id, err := o.storage.OrderDrug().Update(ctx, updateOrderDrug)
	if err != nil {
		fmt.Println("error in servise layer updating order drug  by id", err.Error())
		return models.OrderDrug{}, err
	}

	orderDrug, err := o.storage.OrderDrug().Get(ctx, models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		fmt.Println("error in service layer getting order drug after update", err.Error())
		return models.OrderDrug{}, err
	}

	return orderDrug, nil
}

func (o orderDrugService) Delete(ctx context.Context, id string) error {

	err := o.storage.OrderDrug().Delete(ctx, id)

	return err
}
