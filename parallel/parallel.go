package parallel

import (
	"fmt"
	"testing"
)

func CustomEfficiencyStats(serial func(*testing.B), parallel func(*testing.B, int), threadCounts []int) {
	serialB := testing.Benchmark(serial)

	for _, threadCount := range threadCounts {

		parallelB := testing.Benchmark(func(b *testing.B) {
			parallel(b, threadCount)
		})

		speedup := float32(serialB.NsPerOp()) / float32(parallelB.NsPerOp())
		maxCost := parallelB.NsPerOp() * int64(threadCount) // TODO: ideally use sum thread wall time for this metric

		fmt.Printf("\t[%dt]\tspeedup:%f\tefficiency:%f\tmax total cost:%dns\ttotal overhead time:%dns\n", threadCount, speedup, speedup/float32(threadCount), maxCost, maxCost-serialB.NsPerOp())
	}
}
