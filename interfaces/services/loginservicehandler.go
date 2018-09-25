package services

import (
	"net/http"
	_ "fmt"

	"time"
	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
	usecase "migoV2/usecases/login"
	_ "migoV2/bootstrap"	
)

type BongoUserInteractor interface {
    CreateUser(user usecase.User) (error)
    ValidateUser(user *usecase.User) (error)
}

type LoginServiceHandler struct {
	BongoUserInteractor BongoUserInteractor
}

func (handler *LoginServiceHandler) Signup(c echo.Context) error {
	u := new(usecase.User)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusOK, "Error binding arguments!")
	}

	err := handler.BongoUserInteractor.CreateUser(*u)

	if err != nil {
		return c.JSON(http.StatusOK, "User no registered")
	}

	return  c.JSON(http.StatusOK, u)
}


func (handler *LoginServiceHandler) SignIn(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	
	u := usecase.User{Username:username, Password:password}

	if err := handler.BongoUserInteractor.ValidateUser(&u); err == nil {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = username
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("secret"))

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string {
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}
