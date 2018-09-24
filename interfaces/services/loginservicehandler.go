package services

import (
	"net/http"
	"fmt"

	_ "time"
	_ "github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
	usecase "migoV2/usecases/login"
)

// func (this *CLogin) RegisterServices() {
// 	app.Server.POST("/signup", this.signup)
// 	app.Server.POST("/login", this.signin)
// 	app.Server.GET("/login", this.signin)
// }

type BongoCreateUserInteractor interface {
    CreateUser(user usecase.User) (error)
}

type LoginServiceHandler struct {
	BongoCreateUserInteractor BongoCreateUserInteractor
	Mierda int
}

func (handler *LoginServiceHandler) Signup(c echo.Context) error {
	u := new(usecase.User)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusOK, "Error binding arguments!")
	}

	err := handler.BongoCreateUserInteractor.CreateUser(*u)

	if err != nil {
		fmt.Println(":::", err)
		return c.JSON(http.StatusOK, "User no registered")
	}

	return  c.JSON(http.StatusOK, u)
}


// func (handler *LoginServiceHandler) signin(c Context) error {
// 	username := c.FormValue("username")
// 	password := c.FormValue("password")
	
// 	if err := handler.BongoCreateUserInteractor.Login(username, password); err != nil {
// 		token := jwt.New(jwt.SigningMethodHS256)

// 		// Set claims
// 		claims := token.Claims.(jwt.MapClaims)
// 		claims["name"] = username
// 		claims["admin"] = true
// 		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

// 		// Generate encoded token and send it as response.
// 		t, err := token.SignedString([]byte("secret"))

// 		if err != nil {
// 			return err
// 		}

// 		return c.JSON(http.StatusOK, map[string]string{
// 			"token": t,
// 		})
// 	}

// 	return echo.ErrUnauthorized
// }
