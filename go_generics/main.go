package main

import "fmt"

func main() {
	intValues := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floatValues := map[string]float64{
		"first":  34.90,
		"second": 12.64,
	}

	fmt.Printf("Integers: %v\n", SumInts(intValues))
	fmt.Printf("Floats: %v\n", SumFloats(floatValues))
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
