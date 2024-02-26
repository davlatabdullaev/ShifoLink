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

type ordersService struct {
	storage storage.IStorage
}

func NewOrdersService(storage storage.IStorage) ordersService {
	return ordersService{
		storage: storage,
	}
}

func (o ordersService) Create(ctx context.Context, createOrders models.CreateOrders) (models.Orders, error) {

	pKey, err := o.storage.Orders().Create(ctx, createOrders)
	if err != nil {
		log.Println("error in service layer while creating orders ", err.Error())
		return models.Orders{}, err
	}

	orders, err := o.storage.Orders().Get(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		log.Println("error in service layer get orders by id")
	}

	return orders, nil
}

func (o ordersService) Get(ctx context.Context, pkey models.PrimaryKey) (models.Orders, error) {

	orders, err := o.storage.Orders().Get(ctx, pkey)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting orders by id", err.Error())
			return models.Orders{}, err
		}
	}

	return orders, nil
}

func (o ordersService) GetList(ctx context.Context, request models.GetListRequest) (models.OrdersResponse, error) {

	orders, err := o.storage.Orders().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting orders  list", err.Error())
			return models.OrdersResponse{}, err
		}
	}
	return orders, nil
}

func (o ordersService) Update(ctx context.Context, updateOrders models.UpdateOrders) (models.Orders, error) {

	id, err := o.storage.Orders().Update(ctx, updateOrders)
	if err != nil {
		fmt.Println("error in servise layer updating orders  by id", err.Error())
		return models.Orders{}, err
	}

	orders, err := o.storage.Orders().Get(ctx, models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		fmt.Println("error in service layer getting orders after update", err.Error())
		return models.Orders{}, err
	}

	return orders, nil
}

func (o ordersService) Delete(ctx context.Context, id string) error {

	err := o.storage.Orders().Delete(ctx, id)

	return err
}
