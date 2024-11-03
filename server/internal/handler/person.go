package handler

import (
	"missing-persons-backend/internal/models"
	"missing-persons-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonHandler struct {
	Service * service.PersonService
	ImageService * service.ImageService

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


func (h*PersonHandler) CreateMissingPersons(c*gin.Context) {


	if err := c.Request.ParseMultipartForm(10<<20); err !=nil {
		c.JSON(http.StatusBadRequest,gin.H{

			"error": "Failed to parse multipart data",
			"details": err.Error (),

		})
	}

	var newPerson models.Person

	err := c.ShouldBindJSON(&newPerson)


	if err != nil {

		c.JSON(http.StatusBadRequest ,gin.H{
			"error": "Invalid request data",
			"details": err.Error(),
		})
	}



	file, err := c.FormFile("image")
    if err == nil && file != nil {
        imageURL, err := h.ImageService.SaveImage(file)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to save image",
                "details": err.Error(),
            })
            return
        }
        newPerson.ImageURL = imageURL
    }



	err= h.Service.CreatePerson(newPerson)

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{

			"error":"Failed to create ne person",
			"details":err.Error(),

		})
		return
	}

	c.JSON (http.StatusOK,gin.H{
		"message" : "Person created successfully",
		"data" : "newPerson",
	})





}