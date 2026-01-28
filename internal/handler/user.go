package handler

import (
	"learn-golang-echo/internal/model"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v5"
)

func GetUser(c *echo.Context) error {
	id := c.Param("id")
	idNum, _ := strconv.Atoi(id)
	now := time.Now()

	p := model.User{
		ID:    idNum,
		Name:  "Dary Kusuma Hardian",
		Email: "darykusuma741@gmail.com",
		Date:  now,
	}
	return c.JSON(http.StatusOK, p)
}
