package routes

import (
	"SPMA-APP/controller/item_lwk_controller"
	"SPMA-APP/controller/policy"
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
	PLC := e.Group("/PLC")

	//LWK PENGISIAN ITEM ITEM LWK
	LWK.POST("/ITEM-LWK", item_lwk_controller.Read_EXCEL_Controller)

	//READ ITEM LWK
	LWK.GET("/ITEM-LWK", item_lwk_controller.ReadItemLwk)

	//Edit Item LWK

	//Pengisian Policy (harga)
	PLC.POST("/POLICY", policy.Input_EXCEL_Policy_Controller)

	//Edit Policy (harga)

	//Pengisian Policy (Bonusan)

	//Edit Policy (Bonusan)

	//PENGISIAN DSO

	return e
}
