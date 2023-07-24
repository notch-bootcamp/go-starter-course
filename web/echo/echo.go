package main

import (
	"github.com/labstack/echo/v4"
	"go-starter-course/web/types"
	"net/http"
)

var locationMap = make(map[string]types.Position)

func GetKnownLocation(c echo.Context) error {
	knownLocation := c.Param("knownLocation")
	postition, ok := locationMap[knownLocation]
	if ok {
		return c.JSON(http.StatusOK, postition)
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "Unknown location")
	}
}

func CreateLocation(c echo.Context) error {
	var location types.Position
	if err := c.Bind(&location); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	locationLabel := c.Param("location")
	locationMap[locationLabel] = location
	return c.NoContent(http.StatusCreated)
}

func main() {
	e := echo.New()

	e.GET("/location/:knownLocation", GetKnownLocation)
	e.POST("/location/:location", CreateLocation)

	e.Start(":9093")
}
