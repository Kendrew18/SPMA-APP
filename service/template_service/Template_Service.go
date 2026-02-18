package template_service

import (
	"Template-golang/db"
	"Template-golang/model/request"
	"Template-golang/model/response"
	"fmt"
	"net/http"
	"strconv"
)

func Template_Service(Requests []request.Request_Item_LWK) (response.Response, error) {

	var res response.Response

	factory_code := ""

	con := db.CreateConGorm().Table("ITEM_LWK")

	err := con.Select("factory_code").Where("factory_code = ?", Requests[0].Factory_code).Order("co ASC").Scan(&factory_code).Error

	if factory_code == "" {

		fmt.Println(Requests[0])

		con := db.CreateConGorm().Table("BARANG")

		co := 0

		err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)
		for i := 1; i < 5; i++ {
			fmt.Println("Angka", i)
		}
		Requests[0].Co = co + 1
		Requests[0].Id_tipe_lwk = "LWK-" + strconv.Itoa(Requests[0].Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Requests[0]
			return res, err.Error
		}

		err = con.Select("co", "id_tipe_lwk", "factory_code", "product_name_1", "product_name_2", "qty").Create(&Requests[0])

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Requests[0]
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
