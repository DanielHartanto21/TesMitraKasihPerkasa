package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"tesMitraKasihPerkasa/helper"
	"tesMitraKasihPerkasa/model/domain"
	"tesMitraKasihPerkasa/model/web"
	"tesMitraKasihPerkasa/repository"
)

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImplementation{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

type ProductServiceImplementation struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func (produk ProductServiceImplementation) Insert(ctx context.Context, request web.InputProductRequest) web.ProductResponse {
	err := produk.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := produk.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	product := domain.Product{
		NamaProduct:  request.NamaProduct,
		JenisProduct: request.JenisProduct,
		HargaProduct: request.HargaProduct,
		UpdatedBy:    request.UpdatedBy,
		Stock:        request.Stock,
	}
	product = produk.ProductRepository.Insert(ctx, tx, product)
	return helper.FullProduct(product)
}

func (produk ProductServiceImplementation) ListProduct(ctx context.Context) []web.ProductResponse {
	tx, err := produk.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	product := produk.ProductRepository.ListProduct(ctx, tx)
	return helper.ListProduct(product)
}

func (produk ProductServiceImplementation) EditStock(ctx context.Context, request web.EditStockRequest) {
	tx, err := produk.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	produk.ProductRepository.EditStock(ctx, tx, request)

}
