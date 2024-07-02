package useCase

import "service-hf-product-p5/internal/core/domain/entity/dto"

type ProductUseCase interface {
	SaveProduct(reqProduct dto.RequestProduct) error
	GetProductByID(uuid string) error
	UpdateProductByID(uuid string, product dto.RequestProduct) error
	GetProductByCategory(category string) error
	DeleteProductByID(uuid string) error
}
