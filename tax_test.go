package decimal

import "testing"

func TestNormalWithholdingTax(t *testing.T) {
	price := 49.99
	taxRate := 0.096
	want := 54.78
	wantTax := 4.7990

	got, gotTax := NormalWithholdingTax(price, taxRate)

	t.Logf("price: %#v, taxRate: %#v", price, taxRate)
	t.Logf("fomula: totalAmt = price + (price * taxRate)")
	t.Logf("fomula: totalAmt = %#v + (%#v * %#v)", price, price, taxRate)
	t.Logf("fomula: totalAmt = %#v + (%#v)", got, gotTax)
	t.Logf("fomula: totalAmt (decimal) = %#v", got)
	t.Logf("fomula: totalAmt (%%.3f   ) = %.3f", got)
	t.Logf("fomula: totalAmt (%%.2f   ) = %.2f", got)

	if got != want {
		t.Errorf("Total Amount: want %#v, got %#v\n", want, got)
	}

	if gotTax != wantTax {
		t.Errorf("Tax Amount: want %#v, got %#v\n", wantTax, gotTax)
	}
}

func TestFixedWithholdingTax(t *testing.T) {
	price := 49.99
	taxRate := 0.096
	want := 54.79
	wantTax := 4.79904

	got, gotTax, exact := FixedWithholdingTax(price, taxRate)

	t.Log("price:", price, "taxRate:", taxRate)
	t.Log("fomula: totalAmt = price + (price * taxRate)")
	t.Logf("fomula: totalAmt = %#v + (%#v * %#v)", price, price, taxRate)
	t.Logf("fomula: totalAmt = %#v + (%#v)", got, gotTax)
	t.Logf("fomula: totalAmt (decimal) = %#v", got)
	t.Logf("fomula: totalAmt (%%.3f   ) = %.3f", got)
	t.Logf("fomula: totalAmt (%%.2f   ) = %.2f", got)

	if got != want {
		t.Errorf("Total Amount: want %#v, got %#v\n", want, got)
	}

	if gotTax != wantTax {
		t.Errorf("Tax Amount: want %#v, got %#v\n", wantTax, gotTax)
	}

	if !exact {
		t.Error("the rounded value of total amount with 2 decimal places, which is 54.789040000000001 \nexact is false because the rounding operation was not exact, since the original value had more decimal places than the desired precision")
	}
}
