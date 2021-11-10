package batch

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func worker(db *db, numOps *uint64) {
	for {
		b := db.newBatch()
		db.commit(b)
		b.free()

		atomic.AddUint64(numOps, 1)
		time.Sleep(time.Microsecond)
	}
}

func TestCGo(t *testing.T) {
	batchTest(100, true)
}
func TestGo(t *testing.T) {
	batchTest(100, false)
}

func batchTest(numWorkers int, cgo bool) {
	var numOps uint64
	var groupInGo = !cgo

	db := newDB(groupInGo)
	for i := 0; i < numWorkers; i++ {
		go worker(db, &numOps)
	}

	start := time.Now()
	lastNow := start
	var lastOps uint64

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		if i%20 == 0 {
			fmt.Println("_elapsed____ops/sec")
		}

		now := time.Now()
		elapsed := now.Sub(lastNow)
		ops := atomic.LoadUint64(&numOps)

		fmt.Printf("%8s %10.1f\n",
			time.Duration(time.Since(start).Seconds()+0.5)*time.Second,
			float64(ops-lastOps)/elapsed.Seconds())

		lastNow = now
		lastOps = ops
	}
}
