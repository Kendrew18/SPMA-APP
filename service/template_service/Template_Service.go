package template_service

import (
	"SPMA-APP/db"
	"SPMA-APP/model/request"
	"SPMA-APP/model/response"
	"net/http"
	"strconv"
)

func Template_Service(Requests []request.Request_Item_LWK) (response.Response, error) {

	var res response.Response
	var err error

	factory_code := ""

	stat := 0

	con := db.CreateConGorm().Table("item_lwk")

	for i := 0; i < len(Requests); i++ {
		err = con.Select("factory_code").Where("factory_code = ?", Requests[i].Factory_code).Order("co ASC").Scan(&factory_code).Error
		if factory_code != "" {
			stat = 1
			break
		}
	}

	if stat == 0 {

		con := db.CreateConGorm().Table("item_lwk")

		co := 0

		err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)
		for i := 0; i < len(Requests); i++ {
			Requests[i].Co = co + 1 + i
			Requests[i].Id_tipe_lwk = "LWK-" + strconv.Itoa(Requests[i].Co)
		}

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Requests
			return res, err.Error
		}

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
	} else {
		res.Status = http.StatusNotAcceptable
		res.Message = "ada item yang telah terdaftar dengan factory code yang sama"
		return res, err
	}

	return res, nil
}
