package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {

	int64Map := map[string]int64{"first": 34, "second": 12}

	float64Map := map[string]float64{"first": 34.99, "second": 12.56}

	fmt.Printf("Non-generic sums: %v and %v\n",
		SumInts(int64Map), SumFloats(float64Map))

	// with type
	// fmt.Printf("Generic Sums: %v and %v\n",
	// 	SumIntsOrFloats[string, int64](int64Map),
	// 	SumIntsOrFloats[string, float64](float64Map))

	// without type
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(int64Map),
		SumIntsOrFloats(float64Map))

	// with type constraints in Interface
	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(int64Map),
		SumNumbers(float64Map))
}

func SumInts(m map[string]int64) int64 {
	var sum int64
	for _, value := range m {
		sum += value
	}
	return sum
}

func SumFloats(m map[string]float64) float64 {

	var sum float64
	for _, value := range m {
		sum += value
	}

	return sum
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var sum V
	for _, value := range m {
		sum += value
	}

	return sum
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, value := range m {
		sum += value
	}

	return sum
}
