package simplejson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecimal(t *testing.T) {
	js, err := NewJson([]byte(`{
		"test": {
			"int": 10,
			"float": 5.150,
			"number" : "1.21"
		}
	}`))

	if assert.Nil(t, err) {
		intDecimal, err := js.GetPath("test", "int").Decimal()
		if assert.Nil(t, err) {
			assert.Equal(t, "10", intDecimal.String())
		}

		floatDecimal, err := js.GetPath("test", "float").Decimal()
		if assert.Nil(t, err) {
			assert.Equal(t, "5.15", floatDecimal.String())
		}

		numberDecimal, err := js.GetPath("test", "number").Decimal()
		if assert.Nil(t, err) {
			assert.Equal(t, "1.21", numberDecimal.String())
		}
	}
}
