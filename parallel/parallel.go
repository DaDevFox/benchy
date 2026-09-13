package parallel

import (
	"fmt"
	"testing"
)

func comparisonBenchmark(name string, serial func(*testing.B), parallel func(*testing.B, int), threadCounts []int) {
	// fmt.Printf("%s/Serial\n", name)
	serialB := testing.Benchmark(serial)
	// fmt.Println(serialB.String())

	for _, threadCount := range threadCounts {

		// fmt.Printf("%s/Parallel[threadCount=%d]\n", name, threadCount)
		parallelB := testing.Benchmark(func(b *testing.B) {
			parallel(b, threadCount)
		})
		// fmt.Println(parallelB.String())

		speedup := float32(serialB.NsPerOp()) / float32(parallelB.NsPerOp())
		maxCost := float32(parallelB.NsPerOp() * int64(threadCount)) // TODO: ideally use sum thread wall time for this metric

		fmt.Printf("\t[%dt]\tspeedup:%f\tefficiency:%f\tmax total cost:%fns\ttotal overhead time:%fns\n", threadCount, speedup, speedup/float32(threadCount), maxCost, maxCost-float32(serialB.NsPerOp()))
	}
}
