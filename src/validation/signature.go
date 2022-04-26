package validation

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Rewphg/iambot/src/data"
	"github.com/labstack/echo/v4"
)

func SignatureValidation(header string, body data.EventPost) (error, bool) {

	strBody, err := json.Marshal(body)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()), false
	}

	log.Println((string(strBody)))
	log.Println(strBody)

	decoded, err := base64.StdEncoding.DecodeString(header)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()), false
	}

	hash := hmac.New(sha256.New, []byte(os.Getenv("Channel_Secret")))
	hash.Write([]byte(string(strBody)))

	ans := hmac.Equal(decoded, hash.Sum(nil))

	log.Printf("Header Input : %s, Channel Secret %s \n", header, os.Getenv("Channel_Secret"))
	return nil, ans
}
