package decimal

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func NormalDosage(patientWeight, dosePerKg float64) float64 {
	totalDose := patientWeight * dosePerKg
	fmt.Println("Dosage:", totalDose)
	return totalDose
}

func FixedDosage(patientWeight, dosePerKg float64) (float64, bool) {
	p := decimal.NewFromFloat(patientWeight)
	d := decimal.NewFromFloat(dosePerKg)
	totalDose := p.Mul(d)
	fmt.Println("Dosage:", totalDose.String())
	v, exact := totalDose.Round(3).Float64()
	// the rounded value of totalDose with 3 decimal places, which is 50.466. exact is false because the rounding operation was not exact, since the original value had more decimal places than the desired precision.
	return v, exact
}
