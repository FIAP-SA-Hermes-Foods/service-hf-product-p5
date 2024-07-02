package application

import (
	"context"
	"errors"
	l "service-hf-product-p5/external/logger"
	ps "service-hf-product-p5/external/strings"
	"service-hf-product-p5/internal/core/domain/entity/dto"
	"service-hf-product-p5/internal/core/domain/rpc"
)

type Application interface {
	GetProductByID(msgID string, uuid string) (*dto.OutputProduct, error)
	SaveProduct(msgID string, product dto.RequestProduct) (*dto.OutputProduct, error)
	UpdateProductByID(msgID string, id string, product dto.RequestProduct) (*dto.OutputProduct, error)
	GetProductByCategory(msgID string, category string) ([]dto.OutputProduct, error)
	DeleteProductByID(msgID string, id string) error
}

type application struct {
	ctx              context.Context
	productRPC       rpc.ProductRPC
	productWorkerRPC rpc.ProductWorkerRPC
}

func NewApplication(ctx context.Context, productRPC rpc.ProductRPC, productWorkerRPC rpc.ProductWorkerRPC) Application {
	return application{
		ctx:              ctx,
		productRPC:       productRPC,
		productWorkerRPC: productWorkerRPC,
	}
}

func (app application) GetProductByID(msgID string, uuid string) (*dto.OutputProduct, error) {
	app.setMessageIDCtx(msgID)

	l.Infof(msgID, "GetProductByIDApp: ", " | ", uuid)

	go app.productRPC.GetProductByID(uuid) // pub

	productRpc, err := app.productWorkerRPC.GetProductByID(uuid)

	if err != nil {
		l.Errorf(msgID, "GetProductByIDApp error: ", " | ", err)
		return nil, err
	}

	if productRpc == nil {
		l.Infof(msgID, "GetProductByIDApp output: ", " | ", nil)
		return nil, nil
	}

	product := &dto.OutputProduct{
		UUID:          productRpc.UUID,
		Name:          productRpc.Name,
		Category:      productRpc.Category,
		Image:         productRpc.Image,
		Description:   productRpc.Description,
		Price:         productRpc.Price,
		CreatedAt:     productRpc.CreatedAt,
		DeactivatedAt: productRpc.CreatedAt,
	}

	l.Infof(msgID, "GetProductByIDApp output: ", " | ", product)
	return product, nil
}

func (app application) SaveProduct(msgID string, product dto.RequestProduct) (*dto.OutputProduct, error) {
	app.setMessageIDCtx(msgID)

	l.Infof(msgID, "SaveProductApp: ", " | ", ps.MarshalString(product))

	go app.productRPC.SaveProduct(product) // pub

	pRepo, err := app.productWorkerRPC.SaveProduct(product)

	if err != nil {
		l.Errorf(msgID, "SaveProductApp error: ", " | ", err)
		return nil, err
	}

	if pRepo == nil {
		l.Infof(msgID, "SaveProductApp output: ", " | ", nil)
		return nil, errors.New("is not possible to save product because it's null")
	}

	out := &dto.OutputProduct{
		UUID:          pRepo.UUID,
		Name:          pRepo.Name,
		Category:      pRepo.Category,
		Image:         pRepo.Image,
		Description:   pRepo.Description,
		Price:         pRepo.Price,
		CreatedAt:     pRepo.CreatedAt,
		DeactivatedAt: pRepo.DeactivatedAt,
	}

	l.Infof(msgID, "SaveProductApp output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app application) GetProductByCategory(msgID string, category string) ([]dto.OutputProduct, error) {
	app.setMessageIDCtx(msgID)

	l.Infof(msgID, "GetProductByCategoryApp: ", " | ", category)
	productList := make([]dto.OutputProduct, 0)

	go app.productRPC.GetProductByCategory(category)

	products, err := app.productWorkerRPC.GetProductByCategory(category)

	if err != nil {
		l.Errorf(msgID, "GetProductByCategoryApp error: ", " | ", err)
		return nil, err
	}

	if products == nil {
		l.Infof(msgID, "GetProductByCategoryApp output: ", " | ", nil)
		return nil, nil
	}

	for i := range products {
		product := dto.OutputProduct{
			UUID:          products[i].UUID,
			Name:          products[i].Name,
			Category:      products[i].Category,
			Image:         products[i].Image,
			Description:   products[i].Description,
			Price:         products[i].Price,
			CreatedAt:     products[i].CreatedAt,
			DeactivatedAt: products[i].CreatedAt,
		}
		productList = append(productList, product)
	}

	l.Infof(msgID, "GetProductByCategoryApp output: ", " | ", productList)
	return productList, nil
}

func (app application) UpdateProductByID(msgID string, id string, product dto.RequestProduct) (*dto.OutputProduct, error) {
	app.setMessageIDCtx(msgID)

	l.Infof(msgID, "UpdateProductByIDApp: ", " | ", id, " | ", ps.MarshalString(product))

	go app.productRPC.UpdateProductByID(id, product)

	p, err := app.productWorkerRPC.UpdateProductByID(id, product)

	if err != nil {
		l.Errorf(msgID, "UpdateProductByIDApp error: ", " | ", err)
		return nil, err
	}

	if p == nil {
		productNullErr := errors.New("is not possible to save product because it's null")
		l.Errorf(msgID, "UpdateProductByIDApp output: ", " | ", productNullErr)
		return nil, productNullErr
	}

	out := &dto.OutputProduct{
		UUID:          p.UUID,
		Name:          p.Name,
		Category:      p.Category,
		Image:         p.Image,
		Description:   p.Description,
		Price:         p.Price,
		CreatedAt:     p.CreatedAt,
		DeactivatedAt: p.DeactivatedAt,
	}

	l.Infof(msgID, "UpdateProductByIDApp output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app application) DeleteProductByID(msgID string, id string) error {
	l.Infof(msgID, "DeleteProductByIDApp: ", " | ", id)

	go app.productRPC.DeleteProductByID(id)

	l.Infof(msgID, "DeleteProductByIDApp output: ", " | ", nil)
	return app.productWorkerRPC.DeleteProductByID(id)
}

func (app application) setMessageIDCtx(msgID string) {
	if app.ctx == nil {
		app.ctx = context.WithValue(context.Background(), l.MessageIDKey, msgID)
		return
	}
	app.ctx = context.WithValue(app.ctx, l.MessageIDKey, msgID)
}
