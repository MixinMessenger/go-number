package number

import (
	"github.com/shopspring/decimal"
)

const (
	presentDecimals    = 8
	persistentDecimals = 32
)

var (
	zero       = New(0, 0)
	presentMin = New(1, presentDecimals).Decimal.Round(presentDecimals)
)

type Decimal struct {
	decimal.Decimal
}

func Zero() Decimal {
	return zero
}

func New(value int64, decimals int32) Decimal {
	return Decimal{decimal.New(value, -decimals).Round(persistentDecimals)}
}

func FromString(source string) Decimal {
	d, _ := decimal.NewFromString(source)
	return Decimal{d.Round(persistentDecimals)}
}

func (a Decimal) Add(b Decimal) Decimal {
	return Decimal{a.Decimal.Add(b.Decimal)}
}

func (a Decimal) Sub(b Decimal) Decimal {
	return Decimal{a.Decimal.Sub(b.Decimal)}
}

func (a Decimal) Div(b Decimal) Decimal {
	return Decimal{a.Decimal.DivRound(b.Decimal, persistentDecimals)}
}

func (a Decimal) Divisible(b Decimal) bool {
	if a.Cmp(b) < 0 {
		return false
	}
	div := a.Div(b)
	return div.Floor().Persist() == div.Persist()
}

func (a Decimal) Mul(b Decimal) Decimal {
	return Decimal{a.Decimal.Mul(b.Decimal).Round(persistentDecimals)}
}

func (a Decimal) Neg() Decimal {
	return Decimal{a.Decimal.Neg()}
}

func (a Decimal) Cmp(b Decimal) int {
	return a.Decimal.Cmp(b.Decimal)
}

func (a Decimal) Floor() Decimal {
	return Decimal{a.Decimal.Floor()}
}

func (a Decimal) Ceil() Decimal {
	return Decimal{a.Decimal.Ceil()}
}

func (a Decimal) Round(decimals int32) Decimal {
	return Decimal{a.Decimal.Round(decimals)}
}

func (a Decimal) RoundFloor(decimals int32) Decimal {
	return a.Mul(New(1, -decimals)).Floor().Mul(New(1, decimals))
}

func (a Decimal) RoundCeil(decimals int32) Decimal {
	return a.Mul(New(1, -decimals)).Ceil().Mul(New(1, decimals))
}

func (a Decimal) Equal(b Decimal) bool {
	return a.Decimal.Equal(b.Decimal)
}

func (a Decimal) Persist() string {
	return a.Decimal.String()
}

func (a Decimal) PresentFloor() string {
	return a.RoundFloor(presentDecimals).Persist()
}

func (a Decimal) PresentCeil() string {
	return a.RoundCeil(presentDecimals).Persist()
}

func (a Decimal) Float64() float64 {
	f, _ := a.Decimal.Float64()
	return f
}

func (a Decimal) Exhausted() bool {
	return a.RoundFloor(presentDecimals).LessThan(presentMin)
}
