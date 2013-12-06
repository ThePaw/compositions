// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package compositions

// Some handy functions. 

import (
	"math"
	"math/rand"
)

const iInf int = math.MaxInt32
const inf float64 = math.MaxFloat64

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func imin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func imax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func iSign(a int) int {
	if a == 0 {
		return 0
	}
	if a < 0 {
		return -1
	}
	return 1
}

// See Hubert et al. (2001).
func sign(a float64) int {
	if a == 0 {
		return 0
	}
	if a < 0 {
		return -1
	}
	return 1
}

func iRound(a float64) int {
	return int(math.Floor(a + 0.5))
}

// returns absolute value of an integer
func iAbs(n int) int {
	switch {
	case n < 0:
		return -n
	case n == 0:
		return 0 // return correctly abs(-0)
	}
	return n
}

func cube(x float64) float64 {
	return x * x * x
}

// Uniform random number. 
func unif(low, high int) int {
	return low + rand.Intn(high-low+1)
}

// Uniform random number. 
func unif64(low, high int64) int64 {
	return low + rand.Int63n(high-low+1)
}

// fact returns the factorial of an int
func fact(n int) int {
	if n < 0 {
		panic("factorial not defined for negative numbers")
	}
	if n == 0 {
		return 1
	}
	f := 1
	for i := 0; i < n; i++ {
		f *= f
	}
	return f
}

// natural logarithm
func ln(x float64) float64 {
	return math.Log(x)
}
