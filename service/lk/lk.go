package lk

import (
	"SPMA-APP/db"
	"SPMA-APP/model/request"
	"SPMA-APP/model/response"
	"net/http"
	"strconv"
)

func Lk_Service(Requests []request.Request_LK) (response.Response, error) {
	var res response.Response
	con := db.DB().Table("lk")

	// Check if factory code already exists
	var Id_lk string

	for i := 0; i < len(Requests); i++ {

		err := con.Select("id_lk").Where("id_cust = ? AND nama_cust = ?", Requests[i].Id_cust, Requests[i].Nama_cust).Order("co ASC").Scan(&Id_lk).Error
		if Id_lk != "" && err == nil {
			res.Status = http.StatusNotAcceptable
			res.Message = "ada item yang telah terdaftar dengan factory code yang sama"
			return res, nil
		}

	}

	// Get the latest co value
	var co int

	// Generate co and Id_tipe_lwk
	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Requests
		return res, err.Error
	}

	for i := 0; i < len(Requests); i++ {
		Requests[i].Co = co + 1 + i
		Requests[i].Id_lk = "PLC-" + strconv.Itoa(Requests[i].Co)
	}

	con = db.DB().Table("lk")

	err = con.Select("co", "id_lk", "id_cust", "nama_cust", "rutin", "non_rutin", "super_star", "lmg").Create(&Requests)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Requests
		return res, err.Error
	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = map[string]int64{
			"rows": err.RowsAffected,
		}
	}
	return res, nil
}
