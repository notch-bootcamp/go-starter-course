package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-starter-course/web/types"
	"net/http"
)

var locationMap = make(map[string]types.Position)

func GetKnownLocation(c *gin.Context) {
	knownLocation := c.Param("knownLocation")
	position, ok := locationMap[knownLocation]
	if ok {
		c.JSON(http.StatusOK, position)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Unknown location")})
		return
	}
}

func CreateLocation(c *gin.Context) {
	var location types.Position
	err := c.BindJSON(&location)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	locationLabel := c.Param("location")
	locationMap[locationLabel] = location
	c.Status(http.StatusCreated)
}

func main() {
	r := gin.Default()

	r.GET("/location/:knownLocation", GetKnownLocation)
	r.POST("/location/:location", CreateLocation)

	r.Run(":9092")
}
