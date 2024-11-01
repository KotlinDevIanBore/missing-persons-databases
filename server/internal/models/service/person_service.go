package service

import (
	"missing-persons-backend/internal/models"
	"missing-persons-backend/repository"
)

type PersonService struct {

	repo * repository.PersonRepository
}



func  ValidatePerson (person models.Person) {
	isFirstNameValid :=person.FirstName!= "" && len(person.FirstName )>1 
	isMiddleNameValid := person.MiddleName!= "" && len (person.MiddleName)>1
	isSurnameValid := person.Surname!= ""&&len (person.Surname)>1
	isAgeValid := person.Age >0
	isGenderValid :=person.Gender =="Male" ||person.Gender =="Female" ||person.Gender =="Other"
	islastSeenLocationValid :=  person.LastSeenLocation != ""
	isLastSeenDateValid 



}