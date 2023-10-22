# Demostration of decimal arithmetic in Go

## Starting point

### Normal Version (Using Floating-Point Arithmetic):

```bash
go test -v -run TestNormalFourDigitSum
```

### Fixed Version (Using Decimal Arithmetic):
	it fixed some precision issues due to the limitations of floating-point arithmetic

```bash
go test -v -run TestFixedFourDigitSum
```

## Go rounding behavior

```bash
go test -v -run GoRoundingBehaviorFloat64
```