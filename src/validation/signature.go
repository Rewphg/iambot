package validation

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func SignatureValidation(header string, req echo.Context) (error, bool) {
	defer req.Request().Body.Close()
	body, err := ioutil.ReadAll(req.Request().Body)

	log.Println(body)

	if err != nil {
		log.Println(err)
		return err, false
	}

	decoded, err := base64.StdEncoding.DecodeString(header)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()), false
	}
	hash := hmac.New(sha256.New, []byte("<channel secret>"))
	hash.Write(body)

	ans := hmac.Equal(decoded, hash.Sum(nil))

	log.Printf("Header Input : %s, Channel Secret %s \n", header, os.Getenv("Channel_Secret"))
	return nil, ans
}
