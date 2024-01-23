package main

import (
	"k-avy/urlShrinker/models"
	"k-avy/urlShrinker/pkg/api"

	"github.com/gin-gonic/gin"
)

func main(){
	models.Connect()
	db:=models.DB
	db.AutoMigrate(&models.Datalink{})
	r:=gin.Default()
	r.POST("/urlshortener",api.ShortenURL)
	r.GET("/urlshortener/:short_url", api.Getlink)
	r.Run()
}