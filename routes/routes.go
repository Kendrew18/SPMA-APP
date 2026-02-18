package routes

import (
	"Template-golang/controller/template_controller"
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

	TMP := e.Group("/TMP")

	//NDL
	TMP.GET("/template", template_controller.Read_EXCEL_Controller)

	return e
}
