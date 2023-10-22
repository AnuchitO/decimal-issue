package decimal

import (
	"fmt"
	"testing"
)

func TestNormalFourDigitSum(t *testing.T) {
	NormalFourDigitSum()
}
func TestFixedFourDigitSumm(t *testing.T) {
	FixedFourDigitSum()
}

func TestGoRoundingBehavior(t *testing.T) {
	fmt.Println("FLOAT64 bits")
	GoRoundingBehaviorFloat64()
	fmt.Println()
	fmt.Println("FLOAT32 bits")
	GoRoundingBehaviorFloat32()
}
