package item_lwk_service

import (
	"SPMA-APP/db"
	"SPMA-APP/model/request"
	"SPMA-APP/model/response"
	"net/http"
	"strconv"
)

func Item_Lwk_Service(Requests []request.Request_Item_LWK) (response.Response, error) {
	var res response.Response

	// Buat koneksi database sekali
	con := db.CreateConGorm().Table("item_lwk")

	// Ekstrak factory_codes dari requests untuk query batch
	factoryCodes := make([]string, len(Requests))
	for i, req := range Requests {
		factoryCodes[i] = req.Factory_code
	}

	// Cek apakah ada factory_code yang sudah ada (query tunggal)
	var existingCount int64
	err := con.Where("factory_code IN ?", factoryCodes).Count(&existingCount).Error
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Database error"
		return res, err
	}
	if existingCount > 0 {
		res.Status = http.StatusNotAcceptable
		res.Message = "Ada item yang telah terdaftar dengan factory code yang sama"
		return res, nil
	}

	// Jika tidak ada duplikat, lanjutkan insert
	// Ambil nilai co terbesar
	var maxCo int
	err = con.Select("co").Order("co DESC").Limit(1).Scan(&maxCo).Error
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Database error"
		return res, err
	}

	// Set Co dan Id_tipe_lwk untuk semua requests
	for i := range Requests {
		Requests[i].Co = maxCo + 1 + i
		Requests[i].Id_tipe_lwk = "LWK-" + strconv.Itoa(Requests[i].Co)
	}

	// Insert batch
	err = con.Select("co", "id_tipe_lwk", "factory_code", "product_name_1", "product_name_2", "qty").Create(&Requests).Error
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Gagal insert data"
		return res, err
	}

	// Respons sukses
	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = map[string]int64{
		"rows": int64(len(Requests)), // Asumsi semua berhasil, atau gunakan RowsAffected jika perlu
	}
	return res, nil
}
