package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	intValues := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floatValues := map[string]float64{
		"first":  34.90,
		"second": 12.64,
	}

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](intValues),
		SumIntsOrFloats[string, float64](floatValues),
	)
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}

	return s
}

func SumInts(m map[string]int64) int64 {
	var s int64
	s = 0

	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	s = 0

	for _, v := range m {
		s += v
	}
	return s
}
