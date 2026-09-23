package d22

import "testing"

var benchmarkHit bool

func BenchmarkSetComparison(b *testing.B) {

	const (
		maxNum       int64 = 1_000_000
		elementCount       = 100_000
	)

	// 固定且互不重复的测试数据，范围在 [-maxNum, maxNum] 内
	values := make([]int64, elementCount)
	for i := range values {
		values[i] = int64(i) - elementCount/2
	}

	// 测试构建后实际占用多少堆内存
	// 每一轮都新建一个集合并写入 10 万个数字
	b.Run("Memory/BitMap_BuildAndFill", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			bm := NewBitMap(maxNum)

			for _, num := range values {
				if err := bm.Add(num); err != nil {
					b.Fatal(err)
				}
			}

			// 防止编译器优化掉 bm
			benchmarkHit, _ = bm.Contains(values[0])
		}
	})

	b.Run("Memory/HashSet_BuildAndFill", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			hs := NewHashSet(maxNum)

			for _, num := range values {
				if err := hs.Add(num); err != nil {
					b.Fatal(err)
				}
			}

			benchmarkHit, _ = hs.Contains(values[0])
		}
	})

	// 测试稳定运行中的吞吐：1/3 Add、1/3 Delete、1/3 Contains
	b.Run("Throughput/BitMap_MixedOps", func(b *testing.B) {
		bm := NewBitMap(maxNum)

		// 初始填满，使 Contains、Delete 都是正常路径
		for _, num := range values {
			_ = bm.Add(num)
		}

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			num := values[i%elementCount]

			switch i % 3 {
			case 0:
				_ = bm.Add(num)
			case 1:
				_ = bm.Delete(num)
			case 2:
				benchmarkHit, _ = bm.Contains(num)
			}
		}
	})

	b.Run("Throughput/HashSet_MixedOps", func(b *testing.B) {
		hs := NewHashSet(maxNum)

		for _, num := range values {
			_ = hs.Add(num)
		}

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			num := values[i%elementCount]

			switch i % 3 {
			case 0:
				_ = hs.Add(num)
			case 1:
				_ = hs.Delete(num)
			case 2:
				benchmarkHit, _ = hs.Contains(num)
			}
		}
	})
}
