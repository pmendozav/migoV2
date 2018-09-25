package services

import (
	"net/http"
	_ "fmt"
	"github.com/labstack/echo"
	usecase "migoV2/usecases/common"
	_ "migoV2/bootstrap"	
)

type BongoCommonInteractor interface {
    FindPartner(wildcard string) ([]usecase.Partner, error)
}

type CommonServiceHandler struct {
	BongoCommonInteractor BongoCommonInteractor
}

func (handler *CommonServiceHandler) FindPartnersWithWildcard(c echo.Context) error {
	wildcard := c.Param("wildcard")

	partners, err := handler.BongoCommonInteractor.FindPartner(wildcard)

	if err != nil {
		return c.JSON(http.StatusOK, "Without info")
	}

	return  c.JSON(http.StatusOK, partners)
}
