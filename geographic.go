package decimal

import (
	"math"

	"github.com/shopspring/decimal"
)

func degrees2Radian(deg float64) float64 {
	return deg * (math.Pi / 180)
}

const earthRadius = 6371.0 // Radius of earth in kilometers. Use 3956 for miles

func NormalGeographicHaversine(lat1, lon1, lat2, lon2 float64) float64 {
	// convert decimal degrees to radians
	lat1, lon1, lat2, lon2 = degrees2Radian(lat1), degrees2Radian(lon1), degrees2Radian(lat2), degrees2Radian(lon2)

	// haversine formula
	dlon := lon2 - lon1
	dlat := lat2 - lat1
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	c := 2.0 * math.Asin(math.Sqrt(a))
	return c * earthRadius
}

func FixedGeographicHaversine(lat1, lon1, lat2, lon2 float64) float64 {
	// convert decimal degrees to radians
	lat1, lon1, lat2, lon2 = degrees2Radian(lat1), degrees2Radian(lon1), degrees2Radian(lat2), degrees2Radian(lon2)

	// haversine formula
	dlon := lon2 - lon1
	dlat := lat2 - lat1
	a := math.Pow(math.Sin(dlat/2), 2) + decimal.NewFromFloat(math.Cos(lat1)).Mul(decimal.NewFromFloat(math.Cos(lat2))).Mul(decimal.NewFromFloat(math.Pow(math.Sin(dlon/2), 2))).InexactFloat64()
	c := 2.0 * math.Asin(math.Sqrt(a))
	return c * earthRadius
}
