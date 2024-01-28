package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pedro-hca/go-studies/09-rest-api/db"
	"github.com/pedro-hca/go-studies/09-rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") //localhost 8080
}
