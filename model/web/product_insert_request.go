package web

type InputProductRequest struct {
	NamaProduct  string  `validate:"required"  json:"name"`
	JenisProduct string  `validate:"required"  json:"type"`
	HargaProduct float64 `validate:"required"  json:"price"`
	UpdatedBy    string  `validate:"required"  json:"updatedBy"`
	Stock        int     `validate:"required"  json:"stock"`
}
