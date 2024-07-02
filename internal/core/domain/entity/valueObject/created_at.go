package valueobject

import "time"

type CreatedAt struct {
	Value time.Time `json:"value,omitempty"`
}

var createdAtFormatLayout string = `02-01-2006 15:04:05`

func (c *CreatedAt) Format() string {
	return c.Value.Format(createdAtFormatLayout)
}
