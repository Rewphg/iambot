package api

import (
	"log"
	"net/http"

	"github.com/Rewphg/iambot/src/action"
	"github.com/Rewphg/iambot/src/data"
	"github.com/labstack/echo/v4"
)

func ResLine(c echo.Context) error {

	body := new(data.EventPost)

	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// if err, ans := validation.SignatureValidation(c.Request().Header.Get("x-line-signature"), body); err != nil || !ans {
	// 	log.Println(ans)
	// 	return err
	// }

	UserInfo, err := action.GetUserData(body.Event[0].Source.UserID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Printf("Recieve Message < %s > from User : < %s > \n", body.Event[0].Message.Text, UserInfo.DisplayName)

	if err := TypeRedirector(*body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "OK")
}
