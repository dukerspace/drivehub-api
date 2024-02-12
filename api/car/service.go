package car

import "github.com/dukerspace/drivehub-api/internal/database"

type Service interface {
	GetAll() ([]database.Car, error)
	Create(input inputCar) (*database.Car, error)
	GetByID(id int) (*database.Car, error)
	Update(id int, input inputCar) (*database.Car, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAll() ([]database.Car, error) {
	return s.repository.GetAll()
}

func (s *service) Create(input inputCar) (*database.Car, error) {
	return s.repository.Create(input)
}

func (s *service) GetByID(id int) (*database.Car, error) {
	return s.repository.GetByID(id)
}

func (s *service) Update(id int, input inputCar) (*database.Car, error) {
	return s.repository.Update(id, input)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
