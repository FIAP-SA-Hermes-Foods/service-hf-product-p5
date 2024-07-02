package rpc

import (
	"context"
	"fmt"
	"service-hf-product-p5/internal/core/domain/entity/dto"
	"service-hf-product-p5/internal/core/domain/rpc"
	op "service-hf-product-p5/product_api_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ rpc.ProductWorkerRPC = (*productWorkerRPC)(nil)

type productWorkerRPC struct {
	ctx  context.Context
	host string
	port string
}

func NewProductWorkerRPC(ctx context.Context, host, port string) rpc.ProductWorkerRPC {
	return productWorkerRPC{ctx: ctx, host: host, port: port}
}

func (p productWorkerRPC) SaveProduct(product dto.RequestProduct) (*dto.OutputProduct, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", p.host, p.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	input := op.CreateProductRequest{
		Name:          product.Name,
		Category:      product.Category,
		Image:         product.Image,
		Description:   product.Description,
		Price:         float32(product.Price),
		CreatedAt:     product.CreatedAt,
		DeactivatedAt: product.DeactivatedAt,
	}

	cc := op.NewProductClient(conn)

	resp, err := cc.CreateProduct(p.ctx, &input)

	if err != nil {
		return nil, err
	}

	var out = dto.OutputProduct{
		UUID:          resp.Uuid,
		Name:          resp.Name,
		Category:      resp.Category,
		Image:         resp.Image,
		Description:   resp.Description,
		Price:         float64(resp.Price),
		CreatedAt:     resp.CreatedAt,
		DeactivatedAt: resp.DeactivatedAt,
	}

	return &out, nil
}

func (p productWorkerRPC) UpdateProductByID(id string, product dto.RequestProduct) (*dto.OutputProduct, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", p.host, p.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	input := op.UpdateProductRequest{
		Uuid:          id,
		Name:          product.Name,
		Category:      product.Category,
		Image:         product.Image,
		Description:   product.Description,
		Price:         float32(product.Price),
		CreatedAt:     product.CreatedAt,
		DeactivatedAt: product.DeactivatedAt,
	}

	cc := op.NewProductClient(conn)

	resp, err := cc.UpdateProduct(p.ctx, &input)

	if err != nil {
		return nil, err
	}

	var out = dto.OutputProduct{
		UUID:          resp.Uuid,
		Name:          resp.Name,
		Category:      resp.Category,
		Image:         resp.Image,
		Description:   resp.Description,
		Price:         float64(resp.Price),
		CreatedAt:     resp.CreatedAt,
		DeactivatedAt: resp.DeactivatedAt,
	}

	return &out, nil
}

func (p productWorkerRPC) GetProductByCategory(category string) ([]dto.OutputProduct, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", p.host, p.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	input := op.GetProductByCategoryRequest{
		Category: category,
	}

	cc := op.NewProductClient(conn)

	resp, err := cc.GetProductByCategory(p.ctx, &input)

	if err != nil {
		return nil, err
	}

	out := make([]dto.OutputProduct, 0)

	for i := range resp.Items {
		var outItem = dto.OutputProduct{
			UUID:          resp.Items[i].Uuid,
			Name:          resp.Items[i].Name,
			Category:      resp.Items[i].Category,
			Image:         resp.Items[i].Image,
			Description:   resp.Items[i].Description,
			Price:         float64(resp.Items[i].Price),
			CreatedAt:     resp.Items[i].CreatedAt,
			DeactivatedAt: resp.Items[i].DeactivatedAt,
		}

		out = append(out, outItem)
	}

	return out, nil
}

func (p productWorkerRPC) GetProductByID(uuid string) (*dto.OutputProduct, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", p.host, p.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	input := op.GetProductByIDRequest{
		Uuid: uuid,
	}

	cc := op.NewProductClient(conn)

	resp, err := cc.GetProductByID(p.ctx, &input)

	if err != nil {
		return nil, err
	}

	out := &dto.OutputProduct{
		UUID:          resp.Uuid,
		Name:          resp.Name,
		Category:      resp.Category,
		Image:         resp.Image,
		Description:   resp.Description,
		Price:         float64(resp.Price),
		CreatedAt:     resp.CreatedAt,
		DeactivatedAt: resp.DeactivatedAt,
	}

	return out, nil
}

func (p productWorkerRPC) DeleteProductByID(id string) error {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", p.host, p.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return err
	}

	defer conn.Close()

	input := op.DeleteProductByIDRequest{
		Uuid: id,
	}

	cc := op.NewProductClient(conn)

	if _, err := cc.DeleteProductByID(p.ctx, &input); err != nil {
		return err
	}

	return nil
}
