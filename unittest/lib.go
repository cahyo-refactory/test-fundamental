package unittest

import "math"

//Kubus comment.
type Kubus struct {
	Sisi float64
}

//Volume comment.
func (k Kubus) Volume() float64 {
	return math.Pow(k.Sisi, 3)
}

//Luas comment.
func (k Kubus) Luas() float64 {
	return math.Pow(k.Sisi, 2) * 6
}

//Keliling comment.
func (k Kubus) Keliling() float64 {
	return k.Sisi * 12
}
