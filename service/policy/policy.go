package policy

import (
	"SPMA-APP/db"
	"SPMA-APP/model/request"
	"SPMA-APP/model/response"
	"net/http"
	"strconv"
	"time"
)

func Policy_Service(Requests []request.Request_Policy) (response.Response, error) {
	var res response.Response
	con := db.DB().Table("policy")

	// Check if factory code already exists
	var factory_code string

	for i := 0; i < len(Requests); i++ {

		date, _ := time.Parse("02-01-06", Requests[i].Tanggal_update)
		Requests[i].Tanggal_update = date.Format("2006-01-02")

		err := con.Select("id_policy").Where("depo = ? AND tanggal_update = ?", Requests[i].Depo, Requests[i].Tanggal_update).Order("co ASC").Scan(&factory_code).Error
		if factory_code != "" && err == nil {
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
		Requests[i].Id_policy = "PLC-" + strconv.Itoa(Requests[i].Co)
	}

	con = db.DB().Table("policy")

	err = con.Select("co", "id_policy", "strata", "nama_product", "harga", "depo", "tanggal_update").Create(&Requests)

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
