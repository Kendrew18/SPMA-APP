package routes

import (
	"SPMA-APP/controller/template_controller"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "SPMA APP API is running...")
	})

	LWK := e.Group("/LWK")

	//LWK PENGISIAN ITEM ITEM LWK
	LWK.POST("/ITEM-LWK", template_controller.Read_EXCEL_Controller)

	//Edit Item LWK

	//Pengisian Policy (harga)

	//Edit Policy (harga)

	//Pengisian Policy (Bonusan)

	//Edit Policy (Bonusan)

	//PENGISIAN DSO

	return e
}
