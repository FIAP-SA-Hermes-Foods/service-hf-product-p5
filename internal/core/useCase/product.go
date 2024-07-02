package useCase

import (
	"errors"
	"service-hf-product-p5/internal/core/domain/entity/dto"
	vo "service-hf-product-p5/internal/core/domain/entity/valueObject"
	"service-hf-product-p5/internal/core/domain/useCase"
	"strings"
)

var _ useCase.ProductUseCase = (*productUseCase)(nil)

type productUseCase struct {
}

func NewProductUseCase() productUseCase {
	return productUseCase{}
}

func (p productUseCase) SaveProduct(reqProduct dto.RequestProduct) error {
	product := reqProduct.Product()

	if err := product.Category.Validate(); err != nil {
		return err
	}

	reqProduct.Category = product.Category.Value

	return nil
}

func (p productUseCase) UpdateProductByID(uuid string, reqProduct dto.RequestProduct) error {
	if len(uuid) < 1 {
		return errors.New("the id is not valid for consult")
	}

	product := reqProduct.Product()

	if err := product.Category.Validate(); err != nil {
		return err
	}

	reqProduct.Category = product.Category.Value

	return nil
}

func (p productUseCase) GetProductByID(uuid string) error {
	if len(uuid) < 1 {
		return errors.New("the id is not valid for consult")
	}
	return nil
}

func (p productUseCase) GetProductByCategory(category string) error {
	if len(category) < 1 {
		return errors.New("the category is not valid for consult")
	}

	if _, ok := vo.CategoryMap[strings.ToLower(category)]; !ok {
		return errors.New("category is not valid")
	}

	return nil
}

func (p productUseCase) DeleteProductByID(uuid string) error {
	if len(uuid) < 1 {
		return errors.New("the id is not valid for consult")
	}
	return nil
}
