package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOSTNAME")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_DATABASE"))

	if err != nil {
		panic(err.Error())
	}

	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello world"})
	})
	r.Run("0.0.0.0:" + os.Getenv("PORT"))

	defer db.Close()
}
