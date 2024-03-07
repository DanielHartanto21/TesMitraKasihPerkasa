package repository

import (
	"context"
	"database/sql"
	"tesMitraKasihPerkasa/model/domain"
	"tesMitraKasihPerkasa/model/web"
)

type ProductRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	ListProduct(ctx context.Context, tx *sql.Tx) []domain.Product
	EditStock(ctx context.Context, tx *sql.Tx, request web.EditStockRequest)
}
