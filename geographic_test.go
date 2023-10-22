package decimal

import (
	"math"
	"testing"

	"github.com/shopspring/decimal"
)

func TestNormalGeographicHaversine(t *testing.T) {
	lat1, lon1, lat2, lon2 := 37.12, -86.67, 33.99, -118.40
	want := 2877.9764156252427

	got := NormalGeographicHaversine(lat1, lon1, lat2, lon2)

	t.Logf("lat1: %#v, lon1: %#v, lat2: %#v, lon2: %#v", lat1, lon1, lat2, lon2)
	t.Logf("fomula: totalDistance = 2 * r * arcsin(sqrt(sin((lat2 - lat1)/2)^2 + cos(lat1) * cos(lat2) * sin((lon2 - lon1)/2)^2))")
	t.Logf("fomula: totalDistance = 2 * %#v * arcsin(sqrt(sin((%#v - %#v)/2)^2 + cos(%#v) * cos(%#v) * sin((%#v - %#v)/2)^2))", earthRadius, lat2, lat1, lat1, lat2, lon2, lon1)
	t.Logf("fomula: totalDistance = %#v * arcsin(sqrt(%#v + %#v * %#v * %#v))", 2*earthRadius, math.Pow(math.Sin((lat2-lat1)/2), 2), math.Cos(lat1), math.Cos(lat2), math.Pow(math.Sin((lon2-lon1)/2), 2))
	t.Logf("fomula: totalDistance = %#v * arcsin(sqrt(%#v))", 2*earthRadius, math.Pow(math.Sin((lat2-lat1)/2), 2)+math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin((lon2-lon1)/2), 2))
	t.Logf("fomula: totalDistance = %#v * arcsin(%#v)", 2*earthRadius, math.Sqrt(math.Pow(math.Sin((lat2-lat1)/2), 2)+math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin((lon2-lon1)/2), 2)))
	t.Logf("fomula: totalDistance = %#v * %#v", 2*earthRadius, math.Asin(math.Sqrt(math.Pow(math.Sin((lat2-lat1)/2), 2)+math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin((lon2-lon1)/2), 2))))
	t.Logf("fomula: totalDistance = %#v", got)
	t.Logf("fomula: totalDistance (decimal) = %#v", got)
	t.Logf("fomula: totalDistance (%%.3f   ) = %.3f", got)
	t.Logf("fomula: totalDistance (%%.2f   ) = %.2f", got)

	if got != want {
		t.Errorf("want %#v, got %#v\n", want, got)
	}
}

func TestFixedGeographicHaversine(t *testing.T) {
	lat1, lon1, lat2, lon2 := 37.12, -86.67, 33.99, -118.40
	want := 2877.976415625242

	got := FixedGeographicHaversine(lat1, lon1, lat2, lon2)

	t.Logf("lat1: %#v, lon1: %#v, lat2: %#v, lon2: %#v", lat1, lon1, lat2, lon2)
	t.Logf("fomula: totalDistance = 2 * r * arcsin(sqrt(sin((lat2 - lat1)/2)^2 + cos(lat1) * cos(lat2) * sin((lon2 - lon1)/2)^2))")
	t.Logf("fomula: totalDistance = 2 * %#v * arcsin(sqrt(sin((%#v - %#v)/2)^2 + cos(%#v) * cos(%#v) * sin((%#v - %#v)/2)^2))", earthRadius, lat2, lat1, lat1, lat2, lon2, lon1)
	t.Logf("fomula: totalDistance = %#v * arcsin(sqrt(%#v + %#v * %#v * %#v))", 2*earthRadius, math.Pow(math.Sin((lat2-lat1)/2), 2), math.Cos(lat1), math.Cos(lat2), math.Pow(math.Sin((lon2-lon1)/2), 2))
	t.Logf("fomula: totalDistance = %#v * arcsin(sqrt(%#v))", 2*earthRadius, math.Pow(math.Sin((lat2-lat1)/2), 2)+decimal.NewFromFloat(math.Cos(lat1)).Mul(decimal.NewFromFloat(math.Cos(lat2))).Mul(decimal.NewFromFloat(math.Pow(math.Sin((lon2-lon1)/2), 2))).InexactFloat64())
	t.Logf("fomula: totalDistance = %#v * arcsin(%#v)", 2*earthRadius, math.Sqrt(math.Pow(math.Sin((lat2-lat1)/2), 2)+decimal.NewFromFloat(math.Cos(lat1)).Mul(decimal.NewFromFloat(math.Cos(lat2))).Mul(decimal.NewFromFloat(math.Pow(math.Sin((lon2-lon1)/2), 2))).InexactFloat64()))
	t.Logf("fomula: totalDistance = %#v * %#v", 2*earthRadius, math.Asin(math.Sqrt(math.Pow(math.Sin((lat2-lat1)/2), 2)+decimal.NewFromFloat(math.Cos(lat1)).Mul(decimal.NewFromFloat(math.Cos(lat2))).Mul(decimal.NewFromFloat(math.Pow(math.Sin((lon2-lon1)/2), 2))).InexactFloat64())))
	t.Logf("fomula: totalDistance = %#v", got)
	t.Logf("fomula: totalDistance (decimal) = %#v", got)
	t.Logf("fomula: totalDistance (%%.3f   ) = %.3f", got)
	t.Logf("fomula: totalDistance (%%.2f   ) = %.2f", got)

	if got != want {
		t.Errorf("want %#v, got %#v\n", want, got)
	}
}
