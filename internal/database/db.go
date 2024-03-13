package database

import (
	"context"
	"errors"
)

type Apartment struct {
	apartment_id   string
	apartment_name string
	address        string
	city           string
	state          string
	zip_code       string
}

type Unit struct {
	unit_id        string
	apartment_id   string
	unit_number    string
	unit_type      string
	square_footage string
	rent           float64
	status         string
	tenant_id      string
}

var (
	ErrFetchingApartment = errors.New("failed to fetch apartment by id")
	ErrNotImplemented    = errors.New("not implemented yet")
)

// all logic will be built on top of Service struct
type Service struct {
	Store Store
}

type Store interface {
	GetApartment(context.Context, string) (Apartment, error)
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetApartment(ctx context.Context, id string) (Apartment, error) {
	apartment, err := s.Store.GetApartment(ctx, id)
	if err != nil {
		return Apartment{}, err
	}
	return apartment, ErrFetchingApartment
}

func (s *Service) UpdateApartment(ctx context.Context, apt Apartment) error {
	//
	return ErrNotImplemented
}
