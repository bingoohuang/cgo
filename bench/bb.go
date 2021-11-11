package bench

import (
	"sync"
	"time"
)

type Config struct {
	N          int
	Coroutines int
}

func Bench(fn func(), c *Config) float64 {
	if c == nil {
		c = &Config{}
	}
	if c.N == 0 {
		c.N = 10000
	}

	if c.Coroutines == 0 {
		c.Coroutines = 100
	}

	var wg sync.WaitGroup
	wg.Add(c.Coroutines)
	start := time.Now()

	times := c.N / c.Coroutines
	for i := 0; i < c.Coroutines; i++ {
		curTimes := times
		if i == 0 {
			curTimes = times + c.N%c.Coroutines
		}
		go work(&wg, fn, curTimes)
	}

	wg.Wait()
	cost := time.Since(start)
	return float64(c.N) / cost.Seconds()
}

func work(wg *sync.WaitGroup, fn func(), times int) {
	defer wg.Done()
	for i := 0; i < times; i++ {
		fn()
	}
}
