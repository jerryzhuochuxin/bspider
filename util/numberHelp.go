package util

import (
	"fmt"
	"strconv"
)

func Round2(value float64) float64 {
	floatStr := fmt.Sprintf("%.2f", value)
	rt, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		panic(err)
	}
	return rt
}
