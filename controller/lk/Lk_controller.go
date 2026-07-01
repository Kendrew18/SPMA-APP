package lk

import (
	"SPMA-APP/model/request"
	"SPMA-APP/service/policy"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

func Input_EXCEL_LK_Controller(c echo.Context) error {
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

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Gagal membaca sheet",
		})
	}

	// Ambil data dari Sheet1
	var data []request.Request_LK
	for i, row := range rows {
		// Lewati header jika ada
		if i == 0 {
			continue
		}

		// Pastikan panjang row sesuai
		if len(row) < 5 {
			continue
		}

		strata, _ := strconv.Atoi(row[0]) // konversi string ke int
		harga, _ := strconv.Atoi(row[2])  // konversi string ke int
		data = append(data, request.Request_LK{
			Strata:         strata,
			Nama_product:   row[1],
			Harga:          harga,
			Depo:           row[3],
			Tanggal_update: row[4],
		})
	}

	validRequests := data[1:]

	// fmt.Println(validRequests)

	result, err := policy.Policy_Service(validRequests)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}
