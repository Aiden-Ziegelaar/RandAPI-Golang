package generators

import (
	"math"
	"randAPI/helpers"
	"testing"
)

func TestBoolean (t *testing.T) {
	var rounds = 10000000
	var a, b int
	helpers.RotateSeed()
	for i := 0; i < rounds; i++ {
		if GenerateBoolean().Result {
			a++
		} else {
			b++
		}
	}
	result := float64(a) / float64(a+b)
	if !(0.499 <  result && result < 0.501) {
		t.Errorf("Expected bias to be in range 0.499-0.501, got %f", result)
	}
}

func TestWeightedBooleanLow (t *testing.T) {
	var rounds = 10000000
	var a, b int
	helpers.RotateSeed()
	for i := 0; i < rounds; i++ {
		if GenerateWeightedBoolean(9, 1).Result {
			a++
		} else {
			b++
		}
	}
	result := float64(a) / float64(a+b)
	if !(0.099 <  result && result < 0.101) {
		t.Errorf("Expected bias to be in range 0.099-0.101, got %f", result)
	}
}

func TestWeightedBooleanHigh (t *testing.T) {
	var rounds = 10000000
	var a, b int
	helpers.RotateSeed()
	for i := 0; i < rounds; i++ {
		if GenerateWeightedBoolean(1, 9).Result {
			a++
		} else {
			b++
		}
	}
	result := float64(a) / float64(a+b)
	if !(0.899 <  result && result < 0.901) {
		t.Errorf("Expected bias to be in range 0.899-0.901, got %f", result)
	}
}

func TestInteger (t *testing.T) {
	var rounds int64 = 10000000
	var midpoint = int64(math.MaxInt32/2)
	var output int64 = 0
	var offset int64 = 500000
	helpers.RotateSeed()
	for i := int64(0); i < rounds; i++ {
		output = output + int64(GenerateInteger().Result)
	}
	if !((midpoint - offset)*rounds <  output && output < (midpoint + offset)*rounds) {
		t.Errorf("Expected average to be in range %d-%d, got %d", midpoint-offset, midpoint+offset, output/rounds)
	}
}

func TestIntegerN (t *testing.T) {
	var rounds int64 = 10000000
	var max = int32(1000000)
	var midpoint = int64(max/2)
	var output int64 = 0
	var offset int64 = 100
	helpers.RotateSeed()
	for i := int64(0); i < rounds; i++ {
		output = output + int64(GenerateIntegerN(max).Result)
	}
	if !((midpoint - offset)*rounds <  output && output < (midpoint + offset)*rounds) {
		t.Errorf("Expected average to be in range %d-%d, got %d", midpoint-offset, midpoint+offset, output/rounds)
	}
}