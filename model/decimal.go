package model

import (
	"bytes"
	"strconv"
)

// Decimal is a monetary value returned by the Zoho Payments API.
type Decimal string

func (d *Decimal) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) == 0 || bytes.Equal(data, []byte("null")) {
		return nil
	}
	if data[0] == '"' {
		*d = Decimal(bytes.Trim(data, `"`))
		return nil
	}
	*d = Decimal(data)
	return nil
}

func (d Decimal) String() string { return string(d) }

func (d Decimal) Float64() (float64, error) { return strconv.ParseFloat(string(d), 64) }
