package decimal

import (
	"testing"
)

func TestNormalExchangeUSD2JPY(t *testing.T) {
	usdAmount := 10.3
	exchangeRate := 112.56
	want := 1159.368

	got := NormalExchangeUSD2JPY(usdAmount, exchangeRate)

	if got != want {
		t.Errorf("want %#v, got %#v\n", want, got)
	}
}

func TestFixedExchangeUSD2JPY(t *testing.T) {
	usdAmount := 10.3
	exchangeRate := 112.56
	want := 1159.368

	got, exact := FixedExchangeUSD2JPY(usdAmount, exchangeRate)

	if got != want {
		t.Errorf("want %#v, got %#v\n", want, got)
	}

	if !exact {
		t.Error("the rounded value of jpy with 3 decimal places, which is 1159.368. \nexact is false because the rounding operation was not exact, since the original value had more decimal places than the desired precision")
	}
}
