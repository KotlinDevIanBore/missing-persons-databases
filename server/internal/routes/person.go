package routes

import (
    "github.com/gin-gonic/gin"
	"missing-persons-backend/internal/handler"
)

type Router struct {
    handler *handler.PersonHandler
    engine  *gin.Engine
}

func (r*Router) SetupRoutes () *gin.Engine {


	v1 := r.engine.Group ("/api/v1")

	{

		r.setupPersonRoutes(v1)
	}

	return r.engine

}

func (r*Router) setupPersonRoutes (rg * gin.RouterGroup){


	persons := rg.Group ("/missing-persons")
	{
		persons.GET ("", r.handler.GetMissingPersons)
	}

}