package decimal

import "fmt"

func NormalFourDigitSum() {
	var f float32
	for i := 1; i < 100; i++ {
		f = f + 0.0001
		no := fmt.Sprintf("%6d", i)
		fmt.Println(no, f)
	}
}

func FixedFourDigitSum() {
	var f float32
	var precise float32 = 10000.0
	for i := 1; i < 100; i++ {
		f = (f * precise) + (0.0001 * precise)
		f = f / precise
		no := fmt.Sprintf("%6d", i)
		fmt.Println(no, f)
	}
}

func GoRoundingBehaviorFloat64() {
	var a float64 = 0.02299901
	var b float64 = 0.02299902
	var f float64 = a + b
	fmt.Printf("float64 bit not rounding: %#v\n", f)
	fmt.Printf("float64 bit %%.3f        : %.3f\n", f)
	fmt.Printf("float64 bit %%.4f        : %.4f\n", f)
	fmt.Printf("float64 bit %%.5f        : %.5f\n", f)
}

func GoRoundingBehaviorFloat32() {
	var a float32 = 0.02299901
	var b float32 = 0.02299902
	var f float32 = a + b
	fmt.Printf("float32 bit not rounding: %#v\n", f)
	fmt.Printf("float32 bit %%.3f        : %.3f\n", f)
	fmt.Printf("float32 bit %%.4f        : %.4f\n", f)
	fmt.Printf("float32 bit %%.5f        : %.5f\n", f)
}
