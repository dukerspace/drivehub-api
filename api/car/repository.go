package car

import (
	"errors"

	"github.com/dukerspace/drivehub-api/internal/constant"
	"github.com/dukerspace/drivehub-api/internal/database"
	"github.com/dukerspace/drivehub-api/internal/database/helper"
)

type Repository interface {
	GetAll() ([]database.Car, error)
	Create(input inputCar) (*database.Car, error)
	GetByID(id int) (*database.Car, error)
	Update(id int, input inputCar) (*database.Car, error)
	Delete(id int) error
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]database.Car, error) {
	cars, err := helper.LoadFile(constant.DB_NAME)
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func (repo *repository) Create(input inputCar) (*database.Car, error) {
	data := database.Car{
		Name:     input.Name,
		Price:    input.Price,
		Discount: input.Discount,
	}
	cars, err := helper.LoadFile(constant.DB_NAME)
	if err != nil {
		return nil, err
	}
	cars = append(cars, data)
	helper.SaveToFile(constant.DB_NAME, cars)
	return &data, nil
}

func (repo *repository) GetByID(id int) (*database.Car, error) {
	cars, err := helper.LoadFile(constant.DB_NAME)
	if err != nil {
		return nil, err
	}

	if len(cars) < id+1 {
		return nil, errors.New("Not found")
	}
	car := (cars)[id]
	return &car, nil
}

func (repo *repository) Update(id int, input inputCar) (*database.Car, error) {
	cars, err := helper.LoadFile(constant.DB_NAME)
	if err != nil {
		return nil, err
	}

	if len(cars) < id+1 {
		return nil, errors.New("Not found")
	}
	data := database.Car{
		Name:     input.Name,
		Price:    input.Price,
		Discount: input.Discount,
	}
	(cars)[id] = data
	helper.SaveToFile(constant.DB_NAME, cars)
	return &(cars)[id], nil
}

func (repo *repository) Delete(id int) error {
	cars, err := helper.LoadFile(constant.DB_NAME)
	if err != nil {
		return err
	}

	if len(cars) < id+1 {
		return errors.New("Not found")
	}
	cars = append((cars)[:id], (cars)[id+1:]...)
	helper.SaveToFile(constant.DB_NAME, cars)
	return nil
}
