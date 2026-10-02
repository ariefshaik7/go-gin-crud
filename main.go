package main

import (
	"crud-api/config"
	"crud-api/internal/users"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	uri := os.Getenv("MONGODB_URI")

	DBName := os.Getenv("DB_NAME")

	client := config.ConnectDB(uri)

	defer config.DisconnectDB(client)

	rep := &users.UserRepository{
		DB:     client,
		DBName: DBName,
	}

	ser := &users.UserService{
		UserRepository: rep,
	}

	han := &users.UserHandler{
		UserService: ser,
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	users.RegisterRoutes(r, han)

	r.Run(":8080")

}
