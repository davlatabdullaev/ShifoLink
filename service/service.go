package service

import "shifolink/storage"

type IServiceManager interface {
	Author() authorService
	//other structs

}

type Service struct {
	authorService authorService
	// other structs
}

func New(storage storage.IStorage) Service {
	services := Service{}

	services.authorService = NewAuthorService(storage)
	// other services

	return services
}

func (s Service) Author() authorService {
	return s.authorService
}
