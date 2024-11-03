package handler

import (
	"missing-persons-backend/internal/models"
	"missing-persons-backend/internal/service"
	"net/http"
	"strconv"

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
	return
}


c.JSON(http.StatusOK,persons)

}


func (h*PersonHandler) CreateMissingPersons(c*gin.Context) {


	if err := c.Request.ParseMultipartForm(10<<20); err !=nil {
		c.JSON(http.StatusBadRequest,gin.H{

			"error": "Failed to parse multipart data",
			"details": err.Error (),

		})
		return
	}


	newPerson := models.Person{
		FirstName:        c.PostForm("first_name"),
		MiddleName:       c.PostForm("middle_name"),
		Surname:         c.PostForm("surname"),
		Gender:          c.PostForm("gender"),
		LastSeenLocation: c.PostForm("lastseen_location"),
		LastSeenDate:    c.PostForm("lastseen_date"),
		ContactPerson:    c.PostForm("contact_person"),
		ContactPhone:     c.PostForm("contact_phone"),
		ContactEmail:     c.PostForm("contact_email"),
	}

	if ageStr := c.PostForm("age"); ageStr != "" {
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid age format",
				"details": err.Error(),
			})
			return
		}
		newPerson.Age = age
	}



	file, err := c.FormFile("image_url")
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
		"data" : newPerson,
	})





}