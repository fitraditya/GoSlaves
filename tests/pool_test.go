package slaves

import (
	"testing"

	"github.com/themester/GoSlaves"
)

func BenchmarkSlavePool(b *testing.B) {
	ch := make(chan int, b.N)

	sp := slaves.NewPool(0, func(obj interface{}) {
		ch <- obj.(int)
	})

	go func() {
		for i := 0; i < b.N; i++ {
			sp.Serve(i)
		}
	}()

	var i = 0
	for i < b.N {
		select {
		case <-ch:
			i++
		}
	}

	close(ch)
	sp.Close()
}
