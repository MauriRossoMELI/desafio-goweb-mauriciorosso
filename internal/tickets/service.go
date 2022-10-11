package tickets

import (
	"desafio-go-web/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
	AverageDestination(destination string, hour string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Ticket, error) {
	ticketsList, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ticketsList, nil
}

func (s *service) GetTicketByDestination(destination string) ([]domain.Ticket, error) {
	ticketsList, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return nil, err
	}

	return ticketsList, nil
}

func (s *service) AverageDestination(destination string, hour string) (float64, error) {
	avg, err := s.repository.AverageDestination(destination, hour)
	if err != nil {
		return 0, err
	}

	return avg, nil
}
