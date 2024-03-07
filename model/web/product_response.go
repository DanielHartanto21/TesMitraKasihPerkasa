package web

type ProductResponse struct {
	ProductID    int     `json:"productID"`
	NamaProduct  string  `json:"name"`
	JenisProduct string  `json:"type"`
	HargaProduct float64 `json:"price"`
	UpdatedBy    string  `json:"price"`
	Stock        int     `json:"stock"`
}
