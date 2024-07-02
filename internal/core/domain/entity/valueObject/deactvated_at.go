package valueobject

import (
	"errors"
	"time"
)

type DeactivatedAt struct {
	Value *time.Time `json:"value,omitempty"`
}

var deactivatedAtFormatLayout string = `02-01-2006 15:04:05`

func (d *DeactivatedAt) Format() string {
	if d.Value == nil {
		return "null"
	}

	f := d.Value.Format(deactivatedAtFormatLayout)

	return f
}

var (
	deactivatedAtSaveFromLayout string = `02-01-2006 15:04:05`
	deactivatedAtSaveToLayout   string = `2006-01-02 15:04:05.999999`
)

func (d *DeactivatedAt) SetTimeFromString(du string) error {
	if len(du) == 0 {
		return nil
	}

	if d.Value == nil {
		return errors.New("is not possible set time at deactivatedAt because value is null")
	}

	t, err := time.Parse(deactivatedAtSaveFromLayout, du)

	if err != nil {
		return err
	}

	fmtT := t.Format(deactivatedAtSaveToLayout)

	tt, err := time.Parse(deactivatedAtSaveToLayout, fmtT)

	if err != nil {
		return err
	}

	d.Value = &tt

	return nil
}
