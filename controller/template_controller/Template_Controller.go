package template_controller

import (
	"Template-golang/model/request"
	"Template-golang/service/template_service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Template_Controller(c echo.Context) error {
	var Request request.Request
	Request.Id = c.FormValue("id")
	Request.Nama = c.FormValue("nama")

	result, err := template_service.Template_Service(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
