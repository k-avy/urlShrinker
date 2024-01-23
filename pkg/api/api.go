package api

import (
	"k-avy/urlShrinker/models"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func String(length int) string {
  return StringWithCharset(length, charset)
}

func ShortenURL(c *gin.Context){
   var input models.Postlink

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,err);
	}
	var datalink models.Datalink
	result:= models.DB.Where(models.Datalink{Link:input.Link}).First(&datalink)
	if result.Error!=nil{
		str:=String(32)
		datalink:=models.Datalink{
			Shortlink: str,
			Link: input.Link,
		}
		models.DB.Create(datalink)
		c.JSON(http.StatusCreated, datalink)
	}
	c.JSON(http.StatusCreated, datalink)
}
func Getlink(c *gin.Context){
	short := c.Param("short_url")
	if short==""{
		return
	}
	var url models.Datalink
	err:=models.DB.Where(models.Datalink{Shortlink: short}).First(&url).Error
	if err!=nil{
		c.JSON(http.StatusNotFound, "check the url")
	}
	c.JSON(http.StatusPermanentRedirect,url.Link)
}