package dsp

import (
	"math"
	"testing"

	"github.com/gonutz/check"
)

func TestCopyReturnsANewArray(t *testing.T) {
	a := []FLOAT{1, 2, 3}
	b := Copy(a)
	a[1] = 0
	check.Eq(t, a, []FLOAT{1, 0, 3})
	check.Eq(t, b, []FLOAT{1, 2, 3})
}

func TestCopyingEmptySliceReturnsEmptySlice(t *testing.T) {
	check.Eq(t, len(Copy(nil)), 0)
}

func TestMinMaxReturnsIndicesAndValuesOfFirstExtremes(t *testing.T) {
	a := []FLOAT{3, 2, 1, 3, 1, 4, 4}
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
	a := []FLOAT{5, 0, 5, 9, 5}
	check.Eq(t, MinIndex(a), 1)
	check.Eq(t, MinValue(a), 0)
	check.Eq(t, MaxIndex(a), 3)
	check.Eq(t, MaxValue(a), 9)
}

func TestAverageFilter(t *testing.T) {
	check.Eq(t, AverageFilter([]FLOAT{2, 4, 6, 8}, 2), []FLOAT{3, 5, 7})
	check.Eq(t, AverageFilter([]FLOAT{1, 2, 3, 4, 5}, 3), []FLOAT{2, 3, 4})
	check.Eq(t, AverageFilter([]FLOAT{0, 2, 4, 6, 8, 10}, 4), []FLOAT{3, 5, 7})
}

func TestAverageFilterLeavesAtLeastOneElement(t *testing.T) {
	check.Eq(t, AverageFilter([]FLOAT{1, 2, 3}, 3), []FLOAT{2})
	check.Eq(t, AverageFilter([]FLOAT{1, 2, 3}, 4), []FLOAT{2})
	check.Eq(t, AverageFilter([]FLOAT{1, 2, 3}, 999), []FLOAT{2})

	check.Eq(t, AverageFilter(nil, 0), []FLOAT{0})
	check.Eq(t, AverageFilter(nil, 1), []FLOAT{0})
	check.Eq(t, AverageFilter(nil, 2), []FLOAT{0})
}

func TestAverageFilterOfWidthOneOrLessReturnsCopyOfInput(t *testing.T) {
	for width := 1; width >= -2; width-- {
		a := []FLOAT{1, 2, 3}
		avg := AverageFilter(a, width)
		a[1] = 0
		check.Eq(t, avg, []FLOAT{1, 2, 3})
		check.Eq(t, a, []FLOAT{1, 0, 3})
	}
}

func TestAverage(t *testing.T) {
	check.Eq(t, Average(nil), 0)
	check.Eq(t, Average([]FLOAT{8}), 8)
	check.Eq(t, Average([]FLOAT{1, 2}), 1.5)
	check.Eq(t, Average([]FLOAT{1, 2, 3}), 2)
}

func TestNegation(t *testing.T) {
	check.Eq(t, Negative([]FLOAT{1, -2, 3}), []FLOAT{-1, 2, -3})
}

func TestDerivative(t *testing.T) {
	check.Eq(t, Derivative([]FLOAT{}), []FLOAT{})
	check.Eq(t, Derivative([]FLOAT{1}), []FLOAT{0})
	check.Eq(t, Derivative([]FLOAT{1, 3}), []FLOAT{2})
	check.Eq(t, Derivative([]FLOAT{1, 3, 4}), []FLOAT{2, 1})
	check.Eq(t, Derivative([]FLOAT{1, 3, 4, 2}), []FLOAT{2, 1, -2})
}

func TestNthDerivative(t *testing.T) {
	check.Eq(t, NthDerivative([]FLOAT{1, 3, 4, 2}, -1), []FLOAT{1, 3, 4, 2})
	check.Eq(t, NthDerivative([]FLOAT{1, 3, 4, 2}, 0), []FLOAT{1, 3, 4, 2})
	check.Eq(t, NthDerivative([]FLOAT{1, 3, 4, 2}, 1), []FLOAT{2, 1, -2})
	check.Eq(t, NthDerivative([]FLOAT{1, 3, 4, 2}, 2), []FLOAT{-1, -3})
	check.Eq(t, NthDerivative([]FLOAT{1, 3, 4, 2}, 3), []FLOAT{-2})
	check.Eq(t, NthDerivative([]FLOAT{1, 3, 4, 2}, 4), []FLOAT{0})
	check.Eq(t, NthDerivative([]FLOAT{1, 3, 4, 2}, 5), []FLOAT{0})
}

func TestAddUsesTheLowestCommonElementCount(t *testing.T) {
	check.Eq(t, Add(), nil)
	check.Eq(t, Add(nil, nil), nil)
	check.Eq(t, Add([]FLOAT{1, 2, 3}), []FLOAT{1, 2, 3})
	check.Eq(t, Add([]FLOAT{1, 2, 3}, []FLOAT{4, 7, 9}), []FLOAT{5, 9, 12})
	check.Eq(t, Add([]FLOAT{1, 2}, []FLOAT{4, 7, 9}), []FLOAT{5, 9})
	check.Eq(t, Add([]FLOAT{1, 2, 3}, []FLOAT{4, 7}), []FLOAT{5, 9})
	check.Eq(t, Add([]FLOAT{1}, []FLOAT{2}, []FLOAT{3}), []FLOAT{6})
}

func TestSubUsesTheLowestCommonElementCount(t *testing.T) {
	check.Eq(t, Sub(), nil)
	check.Eq(t, Sub(nil, nil), nil)
	check.Eq(t, Sub([]FLOAT{5, 2, 1}), []FLOAT{5, 2, 1})
	check.Eq(t, Sub([]FLOAT{9, 8, 3}, []FLOAT{4, 7, 9}), []FLOAT{5, 1, -6})
	check.Eq(t, Sub([]FLOAT{9, 8}, []FLOAT{4, 7, 9}), []FLOAT{5, 1})
	check.Eq(t, Sub([]FLOAT{9, 8, 3}, []FLOAT{4, 7}), []FLOAT{5, 1})
	check.Eq(t, Sub([]FLOAT{5}, []FLOAT{1}, []FLOAT{2}), []FLOAT{2})
}

func TestAddOffsetAddsValueToAll(t *testing.T) {
	check.Eq(t, AddOffset(nil, 1), nil)
	check.Eq(t, AddOffset([]FLOAT{2}, 1), []FLOAT{3})
	check.Eq(t, AddOffset([]FLOAT{2, 3, 4}, -1), []FLOAT{1, 2, 3})
}

func TestEveryNthTakesEveryNthElement(t *testing.T) {
	check.Eq(t, EveryNth(nil, 3), nil)
	check.Eq(t, EveryNth([]FLOAT{1}, 3), []FLOAT{1})
	check.Eq(t, EveryNth([]FLOAT{1, 2}, 3), []FLOAT{1})
	check.Eq(t, EveryNth([]FLOAT{1, 2, 3}, 3), []FLOAT{1})
	check.Eq(t, EveryNth([]FLOAT{1, 2, 3, 4}, 3), []FLOAT{1, 4})
	check.Eq(t, EveryNth([]FLOAT{1, 2, 3, 4, 5}, 3), []FLOAT{1, 4})
	check.Eq(t, EveryNth([]FLOAT{1, 2, 3, 4, 5, 6}, 3), []FLOAT{1, 4})
	check.Eq(t, EveryNth([]FLOAT{1, 2, 3, 4, 5, 6, 7}, 3), []FLOAT{1, 4, 7})

	check.Eq(t, EveryNth([]FLOAT{1, 2, 3}, 1), []FLOAT{1, 2, 3})

	check.Eq(t, EveryNth([]FLOAT{1, 2, 3}, 0), nil)
	check.Eq(t, EveryNth([]FLOAT{1, 2, 3}, -1), nil)
}
