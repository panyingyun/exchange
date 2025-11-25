package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUSD2CNY(t *testing.T) {
	assert := assert.New(t)
	type ExchangeTable struct {
		amount float64
		want   float64
	}
	tables := []ExchangeTable{
		{amount: 1.0, want: 7.10},
		{amount: 1.5, want: 10.65},
	}
	for _, c := range tables {
		assert.InDelta(c.want, USD2CNY(c.amount), 0.0001)
	}
}
