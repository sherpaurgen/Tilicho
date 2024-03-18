package service

import (
	"context"
	"errors"

	apartmentModel "github.com/sherpaurgen/Tilicho/internal/property/models"
)

var (
	ErrFetchingApartment = errors.New("failed to fetch apartment by id")
	ErrNotImplemented    = errors.New("not implemented yet")
)

// all business logic will be built on top of Service struct
type Service struct {
	Store Store
}

type Store interface {
	GetApartment(context.Context, string) (apartmentModel.Apartment, error)
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetApartment(ctx context.Context, id string) (apartmentModel.Apartment, error) {
	apartment, err := s.Store.GetApartment(ctx, id)
	if err != nil {
		return apartmentModel.Apartment{}, err
	}
	return apartment, ErrFetchingApartment
}

func (s *Service) UpdateApartment(ctx context.Context, apt apartmentModel.Apartment) error {
	//
	return ErrNotImplemented
}
