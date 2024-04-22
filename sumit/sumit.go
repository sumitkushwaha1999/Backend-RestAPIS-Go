package sumit

import (
	"net/http"

	"app1/requests"
	"github.com/gin-gonic/gin"
)

func InitSumitRoutes(router *gin.RouterGroup) {
	router.POST("/test", handleSumitRoutes)
}
func handleSumitRoutes(c *gin.Context) {
	var req requests.UserData
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Bad requests",
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"status": true,
		"data":   req,
	})
}
