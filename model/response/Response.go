package response

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseExcel struct {
	FACTORY_CODE   string `json:"factory_code"`
	PRODUCT_NAME_1 string `json:"product_name_1"`
	PRODUCT_NAME_2 string `json:"product_name_2"`
	QTY            int    `json:"qty"`
}
