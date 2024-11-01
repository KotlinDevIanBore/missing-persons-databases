package service

import (
	"fmt"
	"missing-persons-backend/internal/models"
	"missing-persons-backend/repository"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type PersonService struct {

	Repo * repository.PersonRepository
}


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


func ValidatePerson(person models.Person) error {
	var validationErrors []string
	phoneRegex := regexp.MustCompile(`^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$`)

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	parsedDate,_ := time.Parse("2006-01=02",person.LastSeenDate )

	isFirstNameValid := person.FirstName != "" && len(person.FirstName) > 1
	isMiddleNameValid := person.MiddleName != "" && len(person.MiddleName) > 1
	isSurnameValid := person.Surname != "" && len(person.Surname) > 1
	isAgeValid := person.Age > 0
	isGenderValid := person.Gender == "Male" || person.Gender == "Female" || person.Gender == "Other"
	isLastSeenLocationValid := person.LastSeenLocation != ""
	isLastSeenDateValid := person.LastSeenDate != "" && !parsedDate.After(time.Now())
	isContactPersonValid := person.ContactPerson != "" 
	isContactPhoneValid := person.ContactPhone != "" && phoneRegex.MatchString(person.ContactPhone)
	isContactEmailValid := person.ContactEmail != "" && emailRegex.MatchString(person.ContactEmail)

	isContactInfoValid:= isContactPhoneValid ||isContactEmailValid

	if !isFirstNameValid {
		validationErrors = append(validationErrors, "First name must be longer than 1 character and not empty.")
	}

	if !isMiddleNameValid {
		validationErrors = append(validationErrors, "Middle name must be longer than 1 character if provided.")
	}

	if !isSurnameValid {
		validationErrors = append(validationErrors, "Surname must be longer than 1 character and not empty.")
	}

	if !isAgeValid {
		validationErrors = append(validationErrors, "Age must be provided and must be a number greater than 0.")
	}

	if !isGenderValid {
		validationErrors = append(validationErrors, "Gender must be 'Male', 'Female', or 'Other'.")
	}

	if !isLastSeenLocationValid {
		validationErrors = append(validationErrors, "Last seen location must not be empty.")
	}

	if !isLastSeenDateValid {
		validationErrors = append(validationErrors, "Last seen date must not be empty.")
	}

	if !isContactPersonValid {
		validationErrors = append(validationErrors, "Contact person must not be empty.")
	}

	if !isContactPhoneValid {
		validationErrors = append(validationErrors, "Contact phone must not be empty.")
	}

	if !isContactEmailValid {
		validationErrors = append(validationErrors, "Contact email must not be empty.")
	}

	if !isContactInfoValid {
		validationErrors = append(validationErrors, "Provide either contact email or contact phone number")


	}

	if len(validationErrors) > 0 {
		return fmt.Errorf("validation errors: %v", validationErrors)
	}

	return nil
}