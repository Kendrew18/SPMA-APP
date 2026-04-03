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
	con := db.DB().Table("item_lwk")

	// Check if factory code already exists
	var factory_code string

	for i := 0; i < len(Requests); i++ {
		err := con.Select("factory_code").Where("factory_code = ?", Requests[i].Factory_code).Order("co ASC").Scan(&factory_code).Error
		if factory_code != "" && err == nil {
			res.Status = http.StatusNotAcceptable
			res.Message = "ada item yang telah terdaftar dengan factory code yang sama"
			return res, nil
		}
	}

	// Get the latest co value
	var co int
	if err := con.Select("co").Order("co DESC").Limit(1).Scan(&co).Error; err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Requests
		return res, err
	}

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
		Requests[i].Id_tipe_lwk = "LWK-" + strconv.Itoa(Requests[i].Co)
	}

	con = db.DB().Table("item_lwk")

	err = con.Select("co", "id_tipe_lwk", "factory_code", "product_name_1", "product_name_2", "qty").Create(&Requests)

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

func Read_Item_Lwk_Service() (response.Response, error) {
	var res response.Response
	var arr_invent []response.Response_Item_LWK

	con := db.DB().Table("item_lwk")

	err := con.Select("co", "id_tipe_lwk", "factory_code", "product_name_1", "product_name_2", "qty").Order("co ASC").Scan(&arr_invent).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_invent
		return res, err
	}

	if arr_invent == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_invent

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = arr_invent
	}

	return res, nil
}
