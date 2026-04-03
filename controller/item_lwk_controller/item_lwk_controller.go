package item_lwk_controller

import (
	"SPMA-APP/model/request"
	"SPMA-APP/service/item_lwk_service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

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
	var data []request.Request_Item_LWK
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
		data = append(data, request.Request_Item_LWK{
			Product_name_1: row[0],
			Product_name_2: row[1],
			Factory_code:   row[2],
			Qty:            QTY,
		})
	}

	validRequests := data[1:]

	result, err := item_lwk_service.Item_Lwk_Service(validRequests)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}

func ReadItemLwk(c echo.Context) error {
	result, err := item_lwk_service.Read_Item_Lwk_Service()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
