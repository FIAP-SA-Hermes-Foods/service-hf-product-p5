package entity

import (
	vo "service-hf-product-p5/internal/core/domain/entity/valueObject"
)

type Product struct {
	UUID          string           `json:"uuid,omitempty"`
	Name          string           `json:"name,omitempty"`
	Category      vo.Category      `json:"category,omitempty"`
	Image         string           `json:"image,omitempty"`
	Description   string           `json:"description,omitempty"`
	Price         float64          `json:"price,omitempty"`
	CreatedAt     vo.CreatedAt     `json:"createdAt,omitempty"`
	DeactivatedAt vo.DeactivatedAt `json:"deactivatedAt,omitempty"`
}
