package controller

import (
	"app1/persons/service"
	"app1/requests"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePerson(c *gin.Context) {
	var req requests.Person
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Bad request",
		})
		return
	}
	err := service.GetPersonServiceInstance().CreatePerson(c, req)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "request failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "sucess",
	})
}
func GetPersonById(c *gin.Context) {
	var req struct {
		id int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Bad request",
		})
		return
	}
	var ids []int
	ids = append(ids, req.id)
	persons, err := service.GetPersonServiceInstance().GetPerson(c, ids)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "request failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "sucess",
		"data":    persons,
	})
}
func GetPerson(c *gin.Context) {
	var req struct {
		Ids []int `json:"ids" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Bad request",
		})
		return
	}
	persons, err := service.GetPersonServiceInstance().GetPerson(c, req.Ids)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "request failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "sucess",
		"data":    persons,
	})
}
func UpdatePerson(c *gin.Context) {

}
func SoftDeletePerson(c *gin.Context) {

}
