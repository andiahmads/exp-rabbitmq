package service

import "math"

type Hitung interface {
	Luas() (float64, error)
}

type Persegi struct {
	Sisi float64
}

func (p Persegi) Luas() (float64, error) {
	return math.Pow(p.Sisi, 2), nil
}
