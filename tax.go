package decimal

import "github.com/shopspring/decimal"

func NormalWithholdingTax(price, taxRate float64) (float64, float64) {
	tax := price * taxRate
	totalAmt := price + tax
	return totalAmt, tax
}

func FixedWithholdingTax(price, taxRate float64) (float64, float64, bool) {
	p := decimal.NewFromFloat(price)
	t := decimal.NewFromFloat(taxRate)
	tax := p.Mul(t)
	totalAmt := p.Add(tax)
	v, exact := totalAmt.Round(2).Float64()
	return v, tax.InexactFloat64(), exact
}
