package routes

import (
	"app1/devesh"
	personRoute "app1/persons/routes"
	"app1/sumit"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	deveshRoutes := router.Group("/devesh")
	sumitRoutes := router.Group("/sumit")
	personRoutes := router.Group("/v1")
	devesh.InitDeveshRoutes(deveshRoutes)
	sumit.InitSumitRoutes(sumitRoutes)
	personRoute.InitPersonRoutes(personRoutes)
}
