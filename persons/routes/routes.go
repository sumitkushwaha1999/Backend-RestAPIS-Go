package routes

import (
	"app1/persons/controller"

	"github.com/gin-gonic/gin"
)

func InitPersonRoutes(router *gin.RouterGroup) {
	router.GET("/persons", controller.GetPerson)          //get allperson
	router.GET("/person/:id", controller.GetPersonById)   //get person
	router.POST("/person", controller.CreatePerson)       //create person
	router.PUT("/person", controller.UpdatePerson)        //update person
	router.DELETE("/person", controller.SoftDeletePerson) //delete person
}
