package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

// var db []string

type DB struct {
	text string
}

var db []DB

type DataRequest struct {
	Text string `json:"text"`
}

func main() {
	// db = make([]string, 0)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))
	r.GET("/", handler)
	r.POST("/send", postHandler)
	r.Run(":" + os.Getenv("PORT"))
}

func handler(c *gin.Context) {
	var s []string
	for _, v := range db {
		s = append(s, v.text)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": s,
	})
}

func postHandler(c *gin.Context) {
	var data DataRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// db = append(db, data.Text)
	var dataBaru = DB{text: data.Text}

	db = append(db, dataBaru)

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil terkirim", "data": data.Text})
}
