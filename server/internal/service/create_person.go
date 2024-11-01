package service

import (
	"fmt"
	"missing-persons-backend/internal/models"

	"github.com/google/uuid"
)

func (s*PersonService) CreatePerson ( person models.Person) error {


	person.ID =uuid.New().String()
	err := ValidatePerson(person)

	if err != nil {


		return fmt.Errorf("validation failed : %w", err)

	}


	err = s.Repo.CreatePerson(person)

	if err != nil {

		return fmt.Errorf("failed to create person : %w",err)
	}

	return nil
}