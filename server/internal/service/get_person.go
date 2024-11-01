package service

import (
	"fmt"
	"missing-persons-backend/internal/models"

)

func (g *PersonService) GetPerson  () ( []models.Person,error ){




	persons,err := g.Repo.GetPerson()


	if err != nil {

		return nil,  fmt.Errorf("issue fetching missing persons from the database : %w", err)
	}

   

return persons,nil
}