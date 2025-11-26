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

func TestEUR2USD(t *testing.T) {
	assert := assert.New(t)
	type ExchangeTable struct {
		amount float64
		want   float64
	}
	tables := []ExchangeTable{
		{amount: 1.0, want: 1.15},
		{amount: 2.0, want: 2.30},
		{amount: 100.0, want: 115.0},
	}
	for _, c := range tables {
		assert.InDelta(c.want, EUR2USD(c.amount), 0.0001)
	}
}

func TestUSD2EUR(t *testing.T) {
	assert := assert.New(t)
	type ExchangeTable struct {
		amount float64
		want   float64
	}
	tables := []ExchangeTable{
		{amount: 1.0, want: 0.87},
		{amount: 2.0, want: 1.74},
		{amount: 100.0, want: 87.0},
	}
	for _, c := range tables {
		assert.InDelta(c.want, USD2EUR(c.amount), 0.0001)
	}
}

func TestCNY2USD(t *testing.T) {
	assert := assert.New(t)
	type ExchangeTable struct {
		amount float64
		want   float64
	}
	tables := []ExchangeTable{
		{amount: 1.0, want: 0.14},
		{amount: 7.10, want: 0.994},
		{amount: 100.0, want: 14.0},
	}
	for _, c := range tables {
		assert.InDelta(c.want, CNY2USD(c.amount), 0.0001)
	}
}

func BenchmarkEUR2USD(b *testing.B) {
	amount := 100.0
	for i := 0; i < b.N; i++ {
		EUR2USD(amount)
	}
}

func BenchmarkUSD2EUR(b *testing.B) {
	amount := 100.0
	for i := 0; i < b.N; i++ {
		USD2EUR(amount)
	}
}

func BenchmarkUSD2CNY(b *testing.B) {
	amount := 100.0
	for i := 0; i < b.N; i++ {
		USD2CNY(amount)
	}
}

func BenchmarkCNY2USD(b *testing.B) {
	amount := 100.0
	for i := 0; i < b.N; i++ {
		CNY2USD(amount)
	}
}
