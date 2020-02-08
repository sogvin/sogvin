package bench

import (
	"testing"

	"github.com/gregoryv/sogvin/internal"
)

func Benchmark_double1(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r, _ = double1(2)
	}
	_ = r
}

func Benchmark_double2(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		double2(&r, 2)
	}
}

func Benchmark_double3(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		double3(&r, 2)
	}
}

func Test_doubles(t *testing.T) {
	var r int
	ok, _k := internal.Assert(t)
	_, err := double1(2)
	ok(err)
	_, err = double1(MAX)
	ok(err)
	_, err = double1(-1)
	_k(err)

	for _, i := range []int{0, 4, MAX} {
		ok(double2(&r, i))
		ok(double3(&r, i))
	}
	_k(double2(&r, -1))
	_k(double3(&r, -1))
	_k(double3(nil, 8))
}
