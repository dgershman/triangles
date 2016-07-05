package main

import "math"

type Float64 float64

const x=math.Pi/180;
func DegToRad(d float64) float64 {
	return d*x
}

func RadToDeg(r float64) float64 {
	return r/x
}

func LegA (b float64, c float64) float64 {
  	return math.Sqrt(math.Pow(c, 2) - math.Pow(b, 2))
}

func LegB (a float64, c float64) float64 {
	return LegA(a, c)
}

func LegC (a float64, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func AngleA (a float64, c float64) float64 {
	return RadToDeg(math.Acos(a / c))
}

func AngleB () float64 {
	return 90;
}

func AngleC (b float64, c float64) float64 {
	return RadToDeg(math.Acos(b / c))
}

func (val Float64) Round(roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * float64(val)
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}