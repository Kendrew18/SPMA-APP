package request

type Request_Policy struct {
	Co             int    `json:"co"`
	Id_policy      string `json:"id_policy"`
	Strata         int    `json:"strata"`
	Nama_product   string `json:"nama_product"`
	Harga          int    `json:"harga"`
	Depo           string `json:"depo"`
	Tanggal_update string `json:"tanggal_update"`
}
