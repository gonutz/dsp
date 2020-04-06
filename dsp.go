package dsp

import "math"

// Copy returns a copy of the given slice.
func Copy(a []FLOAT) []FLOAT {
	c := make([]FLOAT, len(a))
	copy(c, a)
	return c
}

// MinMax returns the indices and values of the minimum and maximum values in
// a. If a is empty, the indices are -1, the minimum is +INF and the maximum is
// is -INF.
func MinMax(a []FLOAT) (minIndex int, minValue FLOAT, maxIndex int, maxValue FLOAT) {
	if len(a) == 0 {
		return -1, FLOAT(math.Inf(1)), -1, FLOAT(math.Inf(-1))
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
func MinIndex(a []FLOAT) int {
	i, _, _, _ := MinMax(a)
	return i
}

// MinValue returns the minimum value in a. If a is empty, +INF is returned.
func MinValue(a []FLOAT) FLOAT {
	_, v, _, _ := MinMax(a)
	return v
}

// MaxIndex returns the index of the maximum value in a. If a is empty, -1 is
// returned.
func MaxIndex(a []FLOAT) int {
	_, _, i, _ := MinMax(a)
	return i
}

// MaxValue returns the minimum value in a. If a is empty, -INF is returned.
func MaxValue(a []FLOAT) FLOAT {
	_, _, _, v := MinMax(a)
	return v
}

// AverageFilter returns a new array of average filtered values over a. The
// resulting array is width-1 smaller than a. Neighboring elements (width
// neighbors) are averaged.
// If the width is 1 or smaller, a copy of the input array is returned.
// If width is greater than len(a), a one-element array with the average value
// over a is returned.
func AverageFilter(a []FLOAT, width int) []FLOAT {
	if width >= len(a) {
		return []FLOAT{Average(a)}
	}

	if width <= 1 {
		return Copy(a)
	}

	b := make([]FLOAT, len(a)-(width-1))
	f := 1.0 / FLOAT(width)

	var slidingSum FLOAT
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
func Average(a []FLOAT) FLOAT {
	if len(a) == 0 {
		return 0
	}

	var sum FLOAT
	for _, v := range a {
		sum += v
	}
	return sum / FLOAT(len(a))
}

// Negative returns a slice of the length of a with all elements the negations
// of those in a.
func Negative(a []FLOAT) []FLOAT {
	n := make([]FLOAT, len(a))
	for i := range n {
		n[i] = -a[i]
	}
	return n
}

// Derivative
func Derivative(a []FLOAT) []FLOAT {
	if len(a) <= 1 {
		return make([]FLOAT, len(a))
	}

	b := make([]FLOAT, len(a)-1)
	for i := range b {
		b[i] = a[i+1] - a[i]
	}
	return b
}

// NthDerivative applies Derivative n times to a. If n is <= 0, a copy of a is
// returned.
func NthDerivative(a []FLOAT, n int) []FLOAT {
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
func Add(a ...[]FLOAT) []FLOAT {
	if len(a) == 0 {
		return nil
	}
	n := len(a[0])
	for _, v := range a {
		if len(v) < n {
			n = len(v)
		}
	}
	sum := make([]FLOAT, n)
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
func Sub(a ...[]FLOAT) []FLOAT {
	if len(a) == 0 {
		return nil
	}
	n := len(a[0])
	for _, v := range a {
		if len(v) < n {
			n = len(v)
		}
	}

	diff := make([]FLOAT, n)
	copy(diff, a[0])

	for i := range diff {
		for j := 1; j < len(a); j++ {
			diff[i] -= a[j][i]
		}
	}
	return diff
}

// AddOffset returns a new array with all values offset greater than in a.
func AddOffset(a []FLOAT, offset FLOAT) []FLOAT {
	b := make([]FLOAT, len(a))
	for i := range b {
		b[i] = a[i] + offset
	}
	return b
}

// EveryNth constructs a new array from every nth item in a. The first item is
// always used. If n is <= 0, an empty array is returned.
func EveryNth(a []FLOAT, n int) []FLOAT {
	if n <= 0 {
		return nil
	}

	b := make([]FLOAT, (len(a)+n-1)/n)
	for i := range b {
		b[i] = a[i*n]
	}
	return b
}

// Repeat makes s slice of FLOAT of length n and sets all values to x. If n <= 0
// the returned slice is empty.
func Repeat(x FLOAT, n int) []FLOAT {
	if n <= 0 {
		return nil
	}
	v := make([]FLOAT, n)
	for i := range v {
		v[i] = x
	}
	return v
}

// Reverse returns a copy of x with elements in reverse order, e.g.
// 1,2,3 -> 3,2,1.
func Reverse(x []FLOAT) []FLOAT {
	y := make([]FLOAT, len(x))
	for i := range y {
		y[i] = x[len(x)-1-i]
	}
	return y
}
