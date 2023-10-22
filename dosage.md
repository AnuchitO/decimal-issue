**Scenario: Pharmaceutical Dosage Calculation**

In the pharmaceutical industry, it's essential to calculate precise dosages of medications, especially for critical patient care. Let's consider a scenario where a healthcare application needs to calculate the correct dosage of a medication for a patient based on their weight and the prescribed dose per kilogram.

**Normal Version (Using Floating-Point Arithmetic):**

```go
package main

import "fmt"

func main() {
    patientWeight := 68.5 // in kilograms
    dosePerKg := 2.5     // mg per kilogram
    totalDose := patientWeight * dosePerKg
    fmt.Printf("Total Dose: %.2f mg\n", totalDose)
}
```

In the normal version, we calculate the total dose by multiplying the patient's weight and the dose per kilogram. However, using floating-point arithmetic can introduce precision issues.

**Fixed Version (Using Decimal Arithmetic):**

```go
package main

import (
    "fmt"
    "github.com/shopspring/decimal"
)

func main() {
    patientWeight := decimal.NewFromFloat(68.5) // in kilograms
    dosePerKg := decimal.NewFromFloat(2.5)     // mg per kilogram
    totalDose := patientWeight.Mul(dosePerKg)
    fmt.Printf("Total Dose: %s mg\n", totalDose.String())
}
```

In the fixed version, we use the "decimal" library to handle patient weight and the dose per kilogram with high precision. This ensures that the total dosage is accurately calculated and rounded to the nearest milligram.

In the pharmaceutical industry, accuracy in medication dosages is critical to patient safety. Using decimal arithmetic can help ensure that medication dosages are calculated precisely, minimizing the risk of over- or under-dosing patients.

Other industries where precise decimal calculations are essential include scientific research, engineering, and manufacturing, especially when dealing with measurements, experiments, and quality control where rounding errors can have significant consequences. In these scenarios, using decimal arithmetic libraries can help maintain data accuracy and integrity.