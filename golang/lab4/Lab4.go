package laba4

import (
	"fmt"
	"math"
)

func Calculate(x float64) float64 {

	return math.Pow(math.Abs(x*x-2.5), 0.25) + math.Pow(math.Log10(x*x), 0.33333333)

}

func laba4() {

	fmt.Println(123)

}
