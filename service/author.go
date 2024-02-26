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

type authorService struct {
	storage storage.IStorage
}

func NewAuthorService(storage storage.IStorage) authorService {
	return authorService{
		storage: storage,
	}
}

func (a authorService) Create(ctx context.Context, createAuthor models.CreateAuthor) (models.Author, error) {

	pKey, err := a.storage.Author().Create(ctx, createAuthor)
	if err != nil {
		log.Println("error in service layer while creating author ", err.Error())
		return models.Author{}, err
	}

	author, err := a.storage.Author().Get(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		log.Println("error in service layer get author by id")
	}

	return author, nil
}

func (a authorService) Get(ctx context.Context, pkey models.PrimaryKey) (models.Author, error) {

	author, err := a.storage.Author().Get(ctx, pkey)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting author by id", err.Error())
			return models.Author{}, err
		}
	}

	return author, nil
}

func (a authorService) GetList(ctx context.Context, request models.GetListRequest) (models.AuthorsResponse, error) {

	authors, err := a.storage.Author().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("error in service layer while getting authors list", err.Error())
			return models.AuthorsResponse{}, err
		}
	}
	return authors, nil
}

func (a authorService) Update(ctx context.Context, updateAuthor models.UpdateAuthor) (models.Author, error) {

	id, err := a.storage.Author().Update(ctx, updateAuthor)
	if err != nil {
		fmt.Println("error in servise layer updating author by id", err.Error())
		return models.Author{}, err
	}

	author, err := a.storage.Author().Get(ctx, models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		fmt.Println("error in service layer getting author after update", err.Error())
		return models.Author{}, err
	}

	return author, nil
}

func (a authorService) Delete(ctx context.Context, id string) error {

	err := a.storage.Author().Delete(ctx, id)

	return err
}

func (a authorService) UpdatePassword(ctx context.Context, request models.UpdateAuthorPassword) error {

	oldPassword, err := a.storage.Author().GetPassword(ctx, request.ID)
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

	if err = a.storage.Author().UpdatePassword(context.Background(), request); err != nil {
		fmt.Println("error in service layer while updating author password ", err.Error())
		return err
	}

	return nil
}
