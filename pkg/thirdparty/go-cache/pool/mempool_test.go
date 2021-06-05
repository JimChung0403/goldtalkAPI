package pool

import "testing"

func Test_SyncPool_AllocSmall(t *testing.T) {
	pool := NewSyncPool(128, 1024, 2)
	mem := pool.Alloc(64)
	if len(mem) != 64 {
		t.Logf("Invalid small alloc:%d", len(mem))
		t.Fail()
	}

	if cap(mem) != 128 {
		t.Logf("Invalid small alloc:%d", len(mem))
		t.Fail()
	}
}

func Test_SyncPool_AllocLarge(t *testing.T) {
	pool := NewSyncPool(128, 1024, 2)
	mem := pool.Alloc(2048)
	if len(mem) != 2048 {
		t.Logf("Invalid small alloc:%d", len(mem))
		t.Fail()
	}

	if cap(mem) != 2048 {
		t.Logf("Invalid small alloc:%d", len(mem))
		t.Fail()
	}
}

func Benchmark_SyncPool_AllocAndFree_128(b *testing.B) {
	pool := NewSyncPool(128, 1024, 2)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pool.Free(pool.Alloc(128))
			b.SetBytes(128)
		}
	})
}

func Benchmark_Normal_Alloc(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := make([]byte, 128)
			if buf == nil {
				b.Error("alloc fail")
				b.FailNow()
			}
			b.SetBytes(int64(len(buf)))
		}
	})
}
