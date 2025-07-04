package service

import (
	"context"

	"github.com/velenac/ordiora/internal/models"
	"github.com/velenac/ordiora/internal/store"
)

type ServiceService struct {
	store    *store.Store
	services *store.ServicesRepository
}

func NewServiceService(s *store.Store) *ServiceService {
	return &ServiceService{
		store:    s,
		services: store.NewServicesRepository(),
	}
}

func (s *ServiceService) GetList(c context.Context) ([]*models.Service, error) {
	data, err := s.services.GetAll(c, s.store.Q())
	if err != nil {
		if err == store.ErrNotFound {
			return nil, ErrNotFound // Return nil if no services found
		}
		return nil, err
	}

	return data, nil
}

func (s *ServiceService) GetByID(c context.Context, id string) (*models.Service, error) {
	data, err := s.services.GetByID(c, s.store.Q(), id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *ServiceService) Create(c context.Context, payload *models.ServicePayload) error {
	newService := &models.Service{
		Description: payload.Description,
		Price:       payload.Price,
	}

	err := s.services.Create(c, s.store.Q(), newService)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceService) Update(c context.Context, id string, payload *models.ServicePayload) error {
	// Check if the service exists
	exist, err := s.services.IsExists(c, s.store.Q(), id)
	if err != nil {
		return err
	}

	if !exist {
		return ErrNotFound
	}

	updateService := &models.Service{
		Description: payload.Description,
		Price:       payload.Price,
	}

	err = s.services.Update(c, s.store.Q(), id, updateService)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceService) Delete(c context.Context, id string) error {
	// Check if the service exists
	exist, err := s.services.IsExists(c, s.store.Q(), id)
	if err != nil {
		return err
	}

	if !exist {
		return ErrNotFound
	}

	err = s.services.Delete(c, s.store.Q(), id)
	if err != nil {
		return err
	}

	return nil
}
