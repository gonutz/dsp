package dsp

import "math"

// Copy returns a copy of the given slice.
func Copy(a []float32) []float32 {
	c := make([]float32, len(a))
	copy(c, a)
	return c
}

// MinMax returns the indices and values of the minimum and maximum values in
// a. If a is empty, the indices are -1, the minimum is +INF and the maximum is
// is -INF.
func MinMax(a []float32) (minIndex int, minValue float32, maxIndex int, maxValue float32) {
	if len(a) == 0 {
		return -1, float32(math.Inf(1)), -1, float32(math.Inf(-1))
	}

	for i := 1; i < len(a); i++ {
		if a[i] < a[minIndex] {
			minIndex = i
		}
		if a[i] > a[maxIndex] {
			maxIndex = i
		}
	}
	minValue = a[minIndex]
	maxValue = a[maxIndex]
	return
}

// MinIndex returns the index of the minimum value in a. If a is empty, -1 is
// returned.
func MinIndex(a []float32) int {
	i, _, _, _ := MinMax(a)
	return i
}

// MinValue returns the minimum value in a. If a is empty, +INF is returned.
func MinValue(a []float32) float32 {
	_, v, _, _ := MinMax(a)
	return v
}

// MaxIndex returns the index of the maximum value in a. If a is empty, -1 is
// returned.
func MaxIndex(a []float32) int {
	_, _, i, _ := MinMax(a)
	return i
}

// MaxValue returns the minimum value in a. If a is empty, -INF is returned.
func MaxValue(a []float32) float32 {
	_, _, _, v := MinMax(a)
	return v
}

// AverageFilter returns a new array of average filtered values over a. The
// resulting array is width-1 smaller than a. Neighboring elements (width
// neighbors) are averaged.
// If the width is 1 or smaller, a copy of the input array is returned.
// If width is greater than len(a), a one-element array with the average value
// over a is returned.
func AverageFilter(a []float32, width int) []float32 {
	if width >= len(a) {
		return []float32{Average(a)}
	}

	if width <= 1 {
		return Copy(a)
	}

	b := make([]float32, len(a)-(width-1))
	f := 1.0 / float32(width)

	var slidingSum float32
	for i := 0; i < width; i++ {
		slidingSum += a[i]
	}
	b[0] = slidingSum * f

	for i := 1; i < len(b); i++ {
		slidingSum += a[i+width-1] - a[i-1]
		b[i] = slidingSum * f
	}

	return b
}

// Average returns the average vaue over a or 0 if a is empty.
func Average(a []float32) float32 {
	if len(a) == 0 {
		return 0
	}

	var sum float32
	for _, v := range a {
		sum += v
	}
	return sum / float32(len(a))
}

// Negative returns a slice of the length of a with all elements the negations
// of those in a.
func Negative(a []float32) []float32 {
	n := make([]float32, len(a))
	for i := range n {
		n[i] = -a[i]
	}
	return n
}

// Derivative returns a slice one item smaller than a, with the differences
// between neighboring items. Result 0 is a[1]-a[0] and so on.
func Derivative(a []float32) []float32 {
	if len(a) <= 1 {
		return make([]float32, len(a))
	}

	b := make([]float32, len(a)-1)
	for i := range b {
		b[i] = a[i+1] - a[i]
	}
	return b
}

// NthDerivative applies Derivative n times to a. If n is <= 0, a copy of a is
// returned.
func NthDerivative(a []float32, n int) []float32 {
	if n <= 0 {
		return Copy(a)
	}
	d := a
	for n > 0 {
		d = Derivative(d)
		n--
	}
	return d
}

// Add returns an array of the sums of the elements in all arrays of a. If the
// arrays in a have different lengths, the smallest of all lengths is used for
// the result.
func Add(a ...[]float32) []float32 {
	if len(a) == 0 {
		return nil
	}
	n := len(a[0])
	for _, v := range a {
		if len(v) < n {
			n = len(v)
		}
	}
	sum := make([]float32, n)
	for i := range sum {
		for j := range a {
			sum[i] += a[j][i]
		}
	}
	return sum
}

// Sub uses the first array in a as the base and subtracts all other arrays from
// it. If the arrays in a have different lengths, the smallest of all lengths is
// used for the result.
func Sub(a ...[]float32) []float32 {
	if len(a) == 0 {
		return nil
	}
	n := len(a[0])
	for _, v := range a {
		if len(v) < n {
			n = len(v)
		}
	}

	diff := make([]float32, n)
	copy(diff, a[0])

	for i := range diff {
		for j := 1; j < len(a); j++ {
			diff[i] -= a[j][i]
		}
	}
	return diff
}

// AddOffset returns a new array with all values offset greater than in a.
func AddOffset(a []float32, offset float32) []float32 {
	b := make([]float32, len(a))
	for i := range b {
		b[i] = a[i] + offset
	}
	return b
}

// EveryNth constructs a new array from every nth item in a. The first item is
// always used. If n is <= 0, an empty array is returned.
func EveryNth(a []float32, n int) []float32 {
	if n <= 0 {
		return nil
	}

	b := make([]float32, (len(a)+n-1)/n)
	for i := range b {
		b[i] = a[i*n]
	}
	return b
}

// Repeat makes s slice of float32 of length n and sets all values to x. If n <= 0
// the returned slice is empty.
func Repeat(x float32, n int) []float32 {
	if n <= 0 {
		return nil
	}
	v := make([]float32, n)
	for i := range v {
		v[i] = x
	}
	return v
}

// Reverse returns a copy of x with elements in reverse order, e.g.
// 1,2,3 -> 3,2,1.
func Reverse(x []float32) []float32 {
	y := make([]float32, len(x))
	for i := range y {
		y[i] = x[len(x)-1-i]
	}
	return y
}

// Scale returns a new array with all values in a scaled by factor.
func Scale(a []float32, factor float32) []float32 {
	b := make([]float32, len(a))
	for i := range b {
		b[i] = a[i] * factor
	}
	return b
}
