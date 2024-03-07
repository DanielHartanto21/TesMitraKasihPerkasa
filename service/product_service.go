package service

import (
	"context"
	"tesMitraKasihPerkasa/model/web"
)

type ProductService interface {
	Insert(ctx context.Context, request web.InputProductRequest) web.ProductResponse
	ListProduct(ctx context.Context) []web.ProductResponse
	EditStock(ctx context.Context, request web.EditStockRequest)
}
