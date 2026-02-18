package template_controller

import (
	"Template-golang/model/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

/*func Template_Controller(c echo.Context) error {
	var Request request.Request
	Request.Id = c.FormValue("id")
	Request.Nama = c.FormValue("nama")

	result, err := template_service.Template_Service(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}*/

func Read_EXCEL_Controller(c echo.Context) error {
	// Buka stream file
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Gagal membuka file",
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Gagal membuka file",
		})
	}
	defer src.Close()

	// Membaca langsung dari io.Reader
	f, err := excelize.OpenReader(src)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Gagal membaca Excel",
		})
	}

	rows, err := f.GetRows("DTBS ITEM")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Gagal membaca sheet",
		})
	}

	// Ambil data dari Sheet1
	var data []response.ResponseExcel
	for i, row := range rows {
		// Lewati header jika ada
		if i == 0 {
			continue
		}

		// Pastikan panjang row sesuai
		if len(row) < 4 {
			continue
		}

		QTY, _ := strconv.Atoi(row[3]) // konversi string ke int
		data = append(data, response.ResponseExcel{
			PRODUCT_NAME_1: row[0],
			PRODUCT_NAME_2: row[1],
			FACTORY_CODE:   row[2],
			QTY:            QTY,
		})
	}

	return c.JSON(http.StatusOK, data)

}
