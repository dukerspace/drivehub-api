package helper

import (
	"encoding/json"
	"io/ioutil"

	"github.com/dukerspace/drivehub-api/internal/database"
)

func LoadFile(filename string) ([]database.Car, error) {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cars []database.Car
	err = json.Unmarshal(fileData, &cars)
	if err != nil {
		return nil, err
	}

	return cars, nil
}

func SaveToFile(filename string, data []database.Car) error {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
