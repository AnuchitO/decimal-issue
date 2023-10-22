package decimal

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func NormalExchangeUSD2JPY(usdAmount, exchangeRate float64) float64 {
	jpyAmount := usdAmount * exchangeRate
	fmt.Println("JPY Amount:", jpyAmount)
	return jpyAmount
}

func FixedExchangeUSD2JPY(usdAmount, exchangeRate float64) (float64, bool) {
	usd := decimal.NewFromFloat(usdAmount)
	exRate := decimal.NewFromFloat(exchangeRate)
	jpy := usd.Mul(exRate)
	fmt.Println("JPY Amount:", jpy.String())
	v, exact := jpy.Round(3).Float64()
	// the rounded value of jpy with 3 decimal places, which is 1159.368. exact is false because the rounding operation was not exact, since the original value had more decimal places than the desired precision.
	return v, exact
}
