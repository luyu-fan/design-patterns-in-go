package __singleton

import (
	"fmt"
	"testing"
)

func TestInstance1(t *testing.T) {
	f1 := Instance("f1")
	t.Logf("Addr: %p\n", f1)
	f2 := Instance("f2")
	t.Logf("Addr: %p\n", f2)
}

func TestInstance2(t *testing.T) {
	i := 0
	for i < 10 {
		// 考虑闭包与并发问题
		j := i
		t.Run(fmt.Sprintf("TestInstance_2_%d", j), func(t *testing.T) {
			// Run方法内部使用Goroutines来运行每一个子用例 阻塞直到这个测试用例成功返回或者是调用了t.Parallel成为并行用例
			// 从执行结果上看不调用t.Parallel可以看作是串行 而调用之后各个测试子用例相互独立 并发执行
			t.Parallel()
			t.Log("[SubTest:]", j)
			f1 := Instance("f1")
			t.Logf("[Addr:] %p\n", f1)
			f2 := Instance("f2")
			t.Logf("[Addr:] %p\n", f2)
		})
		i++
	}
}

func BenchmarkInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f1 := Instance("f1")
		b.Logf("[Addr:] %p\n", f1)
		f2 := Instance("f2")
		b.Logf("[Addr:] %p\n", f2)
	}
}
