package bench

import (
	"runtime"
	"sync"
	"time"
)

type Config struct {
	N            int
	Coroutines   int
	Threads      int
	LockOSThread bool
}

func Bench(fn func(), c *Config) (tps float64) {
	if c == nil {
		c = &Config{}
	}
	if c.N == 0 {
		c.N = 10000
	}
	if c.Coroutines == 0 {
		c.Coroutines = 100
	}
	if c.Threads == 0 {
		c.Threads = int(2.5 * float64(runtime.GOMAXPROCS(0)))
	}
	runtime.GOMAXPROCS(c.Threads)

	var wg sync.WaitGroup
	wg.Add(c.Coroutines)
	start := time.Now()

	times := c.N / c.Coroutines
	work0Times := times + c.N%c.Coroutines
	go c.work(&wg, fn, work0Times)
	for i := 1; i < c.Coroutines; i++ {
		go c.work(&wg, fn, times)
	}

	wg.Wait()
	cost := time.Since(start)
	return float64(c.N) / cost.Seconds()
}

func (c *Config) work(wg *sync.WaitGroup, fn func(), times int) {
	if c.LockOSThread {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
	}

	defer wg.Done()
	for i := 0; i < times; i++ {
		fn()
	}
}
