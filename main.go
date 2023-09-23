package main

import (
	"api/internal/helper/routers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//AMBIL DATA CONNECTION ENVIRONMENT
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	db := initDBConnection()
	fmt.Println("connection success !", db)

	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	//TODO define routes
	api := router.Group("/api/v1")
	route := routers.RoutersInit(api)
	route.ExecRouters(db)

	router.Run(":8080")

}
