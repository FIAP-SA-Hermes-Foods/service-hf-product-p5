package valueobject

import (
	"errors"
	"strings"
)

type Category struct {
	Value string `json:"value,omitempty"`
}

const (
	MealCategory       = "MEAL"
	ComplementCategory = "COMPLEMENT"
	DrinkCategory      = "DRINK"
)

var CategoryMap = map[string]string{
	"meal":           MealCategory,
	"complement":     ComplementCategory,
	"drink":          DrinkCategory,
	"lanche":         MealCategory,
	"acompanhamento": ComplementCategory,
	"bebida":         DrinkCategory,
}

func (v *Category) Validate() error {

	status, ok := CategoryMap[strings.ToLower(v.Value)]

	if !ok {
		return errors.New("category is not valid")
	}

	v.Value = status

	return nil
}


