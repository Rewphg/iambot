package api

import (
	"log"
	"net/http"

	"github.com/Rewphg/iambot/src/action"
	"github.com/Rewphg/iambot/src/data"
	"github.com/Rewphg/iambot/src/validation"
	"github.com/labstack/echo/v4"
)

func ResLine(c echo.Context) error {

	body := new(data.EventPost)

	//bind Body into struct
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//Validator
	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//Validator from line
	if err, ans := validation.SignatureValidation(c.Request().Header.Get("x-line-signature"), c); err != nil || !ans {
		log.Printf("Error Validate, %v, %v \n", err, ans)
	}

	//Get User Data
	UserInfo, err := action.GetUserData(body.Event[0].Source.UserID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Printf("Recieve Message < %s > from User : < %s >, in < %v > , id : < %s >\n", body.Event[0].Message.Text, UserInfo.DisplayName, body.Event[0].Source.Type, UserInfo.UserID)

	if err := TypeRedirector(*body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "OK")
}

func Ping(c echo.Context) error {
	log.Println("Being Ping")

	return c.JSON(http.StatusOK, "OK")
}
