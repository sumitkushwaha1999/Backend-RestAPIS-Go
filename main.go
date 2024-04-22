package main

import (
	"app1/db"
	routes "app1/routers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	//init other needed instances
	if err := db.InitDB(); err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	router := gin.Default()
	routes.InitRoutes(router)
	fmt.Println("listening to port : 8001")
	if err := router.Run(":8001"); err != nil {
		log.Fatal(err)
	}
}
