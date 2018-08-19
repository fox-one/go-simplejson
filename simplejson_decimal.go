package simplejson

import (
	"errors"
	"log"

	"github.com/shopspring/decimal"
)

func (j *Json) Decimal() (decimal.Decimal, error) {
	if s, err := j.String(); err == nil {
		return decimal.NewFromString(s)
	}

	if f, err := j.Float64(); err == nil {
		return decimal.NewFromFloat(f), nil
	}

	return decimal.Zero, errors.New("invalid value type")
}

func (j *Json) MustDecimal(args ...decimal.Decimal) decimal.Decimal {
	var def decimal.Decimal

	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	default:
		log.Panicf("MustMap() received too many arguments %d", len(args))
	}

	d, err := j.Decimal()
	if err == nil {
		return d
	}

	return def
}
