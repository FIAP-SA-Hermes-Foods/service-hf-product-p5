package rpc

import "service-hf-product-p5/internal/core/domain/entity/dto"

type ProductRPC interface {
	GetProductByID(uuid string) (*dto.OutputProduct, error)
	SaveProduct(product dto.RequestProduct) (*dto.OutputProduct, error)
	UpdateProductByID(id string, product dto.RequestProduct) (*dto.OutputProduct, error)
	GetProductByCategory(category string) ([]dto.OutputProduct, error)
	DeleteProductByID(id string) error
}

type ProductWorkerRPC interface {
	GetProductByID(uuid string) (*dto.OutputProduct, error)
	SaveProduct(product dto.RequestProduct) (*dto.OutputProduct, error)
	UpdateProductByID(id string, product dto.RequestProduct) (*dto.OutputProduct, error)
	GetProductByCategory(category string) ([]dto.OutputProduct, error)
	DeleteProductByID(id string) error
}
