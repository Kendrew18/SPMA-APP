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

type Request_Policy_R struct {
	Depo string `json:"depo"`
}

type Request_Policy_Bonus struct {
	Co         int    `json:"co"`
	Id_bonus   string `json:"id_bonus"`
	Nama_bonus string `json:"nama_bonus"`
	Qty        int    `json:"qty"`
	Kode_bonus string `json:"kode_bonus"`
}
