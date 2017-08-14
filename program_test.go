package deferPerf

import "testing"

func Benchmark_executeStandard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		executeStandard()
	}
}

func Benchmark_executeDeferred(b *testing.B) {
	for i := 0; i < b.N; i++ {
		executeDeferred()
	}
}

func Benchmark_executeLongStandard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		executeLongStandard()
	}
}

func Benchmark_executeLongDeferred(b *testing.B) {
	for i := 0; i < b.N; i++ {
		executeLongDeferred()
	}
}
