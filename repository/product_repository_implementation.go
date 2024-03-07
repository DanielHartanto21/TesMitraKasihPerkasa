package repository

import (
	"context"
	"database/sql"
	"tesMitraKasihPerkasa/helper"
	"tesMitraKasihPerkasa/model/domain"
	"tesMitraKasihPerkasa/model/web"
	"time"
)

type ProductRepositoryImplementation struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImplementation{}
}
func (products *ProductRepositoryImplementation) Insert(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "INSERT INTO Product (NamaProduk, JenisProduk, HargaProduk, UpdatedBy, Stock) VALUES ($1, $2, $3, $4, $5)returning ProductID "
	row := tx.QueryRowContext(ctx, SQL, product.NamaProduct, product.JenisProduct, product.HargaProduct, product.UpdatedBy, product.Stock)
	//
	var id int
	err := row.Scan(&id)
	helper.PanicIfError(err)
	//id, _ := result.LastInsertId()
	//helper.PanicIfError(err)
	product.ProductID = id
	return product
}

func (products *ProductRepositoryImplementation) ListProduct(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := "select ProductID,NamaProduk,JenisProduk,HargaProduk,UpdatedBy,Stock from Product"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var product []domain.Product
	for rows.Next() {
		produk := domain.Product{}
		err := rows.Scan(&produk.ProductID, &produk.NamaProduct, &produk.JenisProduct, &produk.HargaProduct, &produk.UpdatedBy, &produk.Stock)
		helper.PanicIfError(err)
		product = append(product, produk)
	}
	return product
}

func (products *ProductRepositoryImplementation) EditStock(ctx context.Context, tx *sql.Tx, request web.EditStockRequest) {
	SQL := "UPDATE Product SET Stock=$1, TanggalUpdate=$2 WHERE ProductID=$3"
	_, err := tx.ExecContext(ctx, SQL, request.Stock, time.Now(), request.ProductID)
	helper.PanicIfError(err)
}
