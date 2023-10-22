package decimal

import "testing"

func TestNormalDosage(t *testing.T) {
	patientWeight := 68.29 // in kilograms
	dosePerKg := 0.739     // mg per kilogram
	want := 50.46631000000001

	got := NormalDosage(patientWeight, dosePerKg)

	t.Logf("patientWeight: %#v, dosePerKg: %#v", patientWeight, dosePerKg)
	t.Logf("fomula: totalDose = patientWeight * dosePerKg")
	t.Logf("fomula: totalDose = %#v * %#v", patientWeight, dosePerKg)
	t.Logf("fomula: totalDose = %#v", got)
	t.Logf("fomula: totalDose (decimal) = %#v", got)
	t.Logf("fomula: totalDose (%%.3f   ) = %.3f", got)
	t.Logf("fomula: totalDose (%%.2f   ) = %.2f", got)

	if got != want {
		t.Errorf("want %#v, got %#v\n", want, got)
	}
}

func TestFixedDosage(t *testing.T) {
	patientWeight := 68.29 // in kilograms
	dosePerKg := 0.739     // mg per kilogram
	want := 50.466

	got, exact := FixedDosage(patientWeight, dosePerKg)

	t.Logf("patientWeight: %#v, dosePerKg: %#v", patientWeight, dosePerKg)
	t.Logf("fomula: totalDose = patientWeight * dosePerKg")
	t.Logf("fomula: totalDose = %#v * %#v", patientWeight, dosePerKg)
	t.Logf("fomula: totalDose = %#v", got)
	t.Logf("fomula: totalDose (decimal) = %#v", got)
	t.Logf("fomula: totalDose (%%.3f   ) = %.3f", got)
	t.Logf("fomula: totalDose (%%.2f   ) = %.2f", got)

	if got != want {
		t.Errorf("want %#v, got %#v\n", want, got)
	}

	if !exact {
		t.Error("the rounded value of total dose with 3 decimal places, which is 50.466. \nexact is true because the rounding operation was exact, since the original value had less decimal places than the desired precision")
	}
}
