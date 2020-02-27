package dsp

import (
	"math"
	"testing"

	"github.com/gonutz/check"
)

func TestCopyReturnsANewArray(t *testing.T) {
	a := []float32{1, 2, 3}
	b := Copy(a)
	a[1] = 0
	check.Eq(t, a, []float32{1, 0, 3})
	check.Eq(t, b, []float32{1, 2, 3})
}

func TestCopyingEmptySliceReturnsEmptySlice(t *testing.T) {
	check.Eq(t, len(Copy(nil)), 0)
}

func TestMinMaxReturnsIndicesAndValuesOfFirstExtremes(t *testing.T) {
	a := []float32{3, 2, 1, 3, 1, 4, 4}
	minIndex, minValue, maxIndex, maxValue := MinMax(a)
	check.Eq(t, minIndex, 2)
	check.Eq(t, minValue, 1)
	check.Eq(t, maxIndex, 5)
	check.Eq(t, maxValue, 4)
}

func TestMinMaxOfEmptyArrayIsNegativeAndInfinite(t *testing.T) {
	minIndex, minValue, maxIndex, maxValue := MinMax(nil)
	check.Eq(t, minIndex, -1)
	check.Eq(t, math.IsInf(float64(minValue), 1), true)
	check.Eq(t, maxIndex, -1)
	check.Eq(t, math.IsInf(float64(maxValue), -1), true)
}

func TestMinMaxIndicesAndValuesCanBeComputedOnTheirOwn(t *testing.T) {
	a := []float32{5, 0, 5, 9, 5}
	check.Eq(t, MinIndex(a), 1)
	check.Eq(t, MinValue(a), 0)
	check.Eq(t, MaxIndex(a), 3)
	check.Eq(t, MaxValue(a), 9)
}

func TestAverageFilter(t *testing.T) {
	check.Eq(t, AverageFilter([]float32{2, 4, 6, 8}, 2), []float32{3, 5, 7})
	check.Eq(t, AverageFilter([]float32{1, 2, 3, 4, 5}, 3), []float32{2, 3, 4})
	check.Eq(t, AverageFilter([]float32{0, 2, 4, 6, 8, 10}, 4), []float32{3, 5, 7})
}

func TestAverageFilterLeavesAtLeastOneElement(t *testing.T) {
	check.Eq(t, AverageFilter([]float32{1, 2, 3}, 3), []float32{2})
	check.Eq(t, AverageFilter([]float32{1, 2, 3}, 4), []float32{2})
	check.Eq(t, AverageFilter([]float32{1, 2, 3}, 999), []float32{2})

	check.Eq(t, AverageFilter(nil, 0), []float32{0})
	check.Eq(t, AverageFilter(nil, 1), []float32{0})
	check.Eq(t, AverageFilter(nil, 2), []float32{0})
}

func TestAverageFilterOfWidthOneOrLessReturnsCopyOfInput(t *testing.T) {
	for width := 1; width >= -2; width-- {
		a := []float32{1, 2, 3}
		avg := AverageFilter(a, width)
		a[1] = 0
		check.Eq(t, avg, []float32{1, 2, 3})
		check.Eq(t, a, []float32{1, 0, 3})
	}
}

func TestAverage(t *testing.T) {
	check.Eq(t, Average(nil), 0)
	check.Eq(t, Average([]float32{8}), 8)
	check.Eq(t, Average([]float32{1, 2}), 1.5)
	check.Eq(t, Average([]float32{1, 2, 3}), 2)
}

func TestNegation(t *testing.T) {
	check.Eq(t, Negative([]float32{1, -2, 3}), []float32{-1, 2, -3})
}

func TestDerivative(t *testing.T) {
	check.Eq(t, Derivative([]float32{}), []float32{})
	check.Eq(t, Derivative([]float32{1}), []float32{0})
	check.Eq(t, Derivative([]float32{1, 3}), []float32{2})
	check.Eq(t, Derivative([]float32{1, 3, 4}), []float32{2, 1})
	check.Eq(t, Derivative([]float32{1, 3, 4, 2}), []float32{2, 1, -2})
}

func TestNthDerivative(t *testing.T) {
	check.Eq(t, NthDerivative([]float32{1, 3, 4, 2}, -1), []float32{1, 3, 4, 2})
	check.Eq(t, NthDerivative([]float32{1, 3, 4, 2}, 0), []float32{1, 3, 4, 2})
	check.Eq(t, NthDerivative([]float32{1, 3, 4, 2}, 1), []float32{2, 1, -2})
	check.Eq(t, NthDerivative([]float32{1, 3, 4, 2}, 2), []float32{-1, -3})
	check.Eq(t, NthDerivative([]float32{1, 3, 4, 2}, 3), []float32{-2})
	check.Eq(t, NthDerivative([]float32{1, 3, 4, 2}, 4), []float32{0})
	check.Eq(t, NthDerivative([]float32{1, 3, 4, 2}, 5), []float32{0})
}

func TestAddUsesTheLowestCommonElementCount(t *testing.T) {
	check.Eq(t, Add(), nil)
	check.Eq(t, Add(nil, nil), nil)
	check.Eq(t, Add([]float32{1, 2, 3}), []float32{1, 2, 3})
	check.Eq(t, Add([]float32{1, 2, 3}, []float32{4, 7, 9}), []float32{5, 9, 12})
	check.Eq(t, Add([]float32{1, 2}, []float32{4, 7, 9}), []float32{5, 9})
	check.Eq(t, Add([]float32{1, 2, 3}, []float32{4, 7}), []float32{5, 9})
	check.Eq(t, Add([]float32{1}, []float32{2}, []float32{3}), []float32{6})
}

func TestSubUsesTheLowestCommonElementCount(t *testing.T) {
	check.Eq(t, Sub(), nil)
	check.Eq(t, Sub(nil, nil), nil)
	check.Eq(t, Sub([]float32{5, 2, 1}), []float32{5, 2, 1})
	check.Eq(t, Sub([]float32{9, 8, 3}, []float32{4, 7, 9}), []float32{5, 1, -6})
	check.Eq(t, Sub([]float32{9, 8}, []float32{4, 7, 9}), []float32{5, 1})
	check.Eq(t, Sub([]float32{9, 8, 3}, []float32{4, 7}), []float32{5, 1})
	check.Eq(t, Sub([]float32{5}, []float32{1}, []float32{2}), []float32{2})
}