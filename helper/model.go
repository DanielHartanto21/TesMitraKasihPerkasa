package helper

import (
	"tesMitraKasihPerkasa/model/domain"
	"tesMitraKasihPerkasa/model/web"
)

func ToLogin(user domain.Users) web.LoginResponse {
	return web.LoginResponse{
		Uuid: user.Uuid,
		Name: user.NamaUser,
	}
}
func FullProduct(prod domain.Product) web.ProductResponse {
	return web.ProductResponse{
		ProductID:    prod.ProductID,
		NamaProduct:  prod.NamaProduct,
		JenisProduct: prod.JenisProduct,
		HargaProduct: prod.HargaProduct,
		UpdatedBy:    prod.UpdatedBy,
		Stock:        prod.Stock,
	}
}
func ListProduct(prod []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range prod {
		productResponses = append(productResponses, FullProduct(product))
	}
	return productResponses
}
