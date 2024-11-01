package handler

import (
	"missing-persons-backend/internal/models"
	"missing-persons-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonHandler struct {
	Service * service.PersonService

}


func (h*PersonHandler) GetMissingPersons( c*gin.Context) {


persons, err:= h.Service.GetPerson();

if err!=nil  {

c.JSON (http.StatusInternalServerError,gin.H{
	"error":"Failed to retrieve missing persons",
	"details": err.Error(),

})
return
}

if len(persons) ==0 {
	c.JSON (http.StatusOK, []models.Person{})
}


c.JSON(http.StatusOK,persons)

}


