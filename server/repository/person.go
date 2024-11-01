package repository

import (
	"database/sql"
	"missing-persons-backend/internal/models"
)

type PersonRepository struct {
	db *sql.DB
}

func (r *PersonRepository) CreatePerson(ID string,
	FirstName string,
	MiddleName string,
	Surname string,
	name string,
	Age int,
	Gender string,
	LastSeenLocation string,
	LastSeenDate string,
	ContactPerson string,
	ContactPhone string,
	ContactEmail string) error {

	values := []interface{}{FirstName, MiddleName, Surname, Age, Gender, LastSeenLocation, LastSeenDate, ContactPerson, ContactPhone, ContactEmail}

	const query = `
	INSERT INTO missing_persons.missing_persons (
	first_name,
	middle_name,
	surname,
	age,
	gender,
	lastseen_location,
	lastseen_date,
	contact_person,
	contact_phone,
	contact_email
	
	

	
	) VALUES (?,?,?,?,?,?,?,?,?.?)


	
	
	`

	_, err := r.db.Exec(query, values)

	return err

}

func (r *PersonRepository) GetPerson() ([]models.Person,error) {

	const query = ` 

	SELECT 
    mst.id AS id,
    mst.first_name AS first_name,
    mst.middle_name AS middle_name,
    mst.surname AS surname,
    mst.age AS age,
    mst.gender AS gender,
    mst.lastseen_location AS lastseen_location,
    mst.lastseen_date AS lastseen_date,
    mst.contact_person AS contact_person,
    mst.contact_phone AS contact_phone,
    mst.contact_email AS contact_email
FROM 
    missing_persons.missing_persons mst;

	
	`

	rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }

	defer rows.Close()

    

	var persons []models.Person


	for rows.Next() {

		var person models.Person
	err:=rows.Scan(

			&person.ID,
			&person.FirstName,
			&person.MiddleName,
			&person.Surname,
			&person.Age,
			&person.Gender,
			&person.LastSeenLocation,
			&person.LastSeenDate,
			&person.ContactPerson,
			&person.ContactPhone,
			&person.ContactEmail,
		);

		if err !=nil {
			return nil ,err
		}



		persons = append(persons, person)

	}

	if err = rows.Err() ;err != nil {
		return nil,err
	}

	return persons, nil

	

}