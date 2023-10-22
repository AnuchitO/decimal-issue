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

## Currency exchange

```bash
go test -v -run TestNormalExchangeUSD2JPY
```

```bash
go test -v -run TestFixedExchangeUSD2JPY
```

## Sales tax calculation

```bash
go test -v -run TestNormalWithholdingTax
```

```bash
go test -v -run TestFixedWithholdingTax
```

## Pharmaceutical dosage calculation

```bash
go test -v -run TestNormalDosage
```

```bash
go test -v -run TestFixedDosage
```

## Geographic distance calculation

```bash
go test -v -run TestNormalGeographicHaversine
```

```bash
go test -v -run TestFixedGeographicHaversine
```