package auth

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	"github.com/martinsd3v/planets/core/tools/communication"
	"github.com/martinsd3v/planets/core/tools/providers/jwt"
)

//Authorize ...
func Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := extractToken(c.Request())
		jwtProvider := jwt.New(viper.GetString("jwt.secretKey"))
		_, err := jwtProvider.CheckToken(token)
		if err != nil {
			comm := communication.New()
			resp := comm.Response(403, "authenticate_failed")
			c.JSON(resp.Status, resp)
			return nil
		}
		return next(c)
	}
}

//get the token from the request body
func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	simpleToken := r.Header.Get("token")
	if simpleToken != "" {
		return simpleToken
	}
	return ""
}
