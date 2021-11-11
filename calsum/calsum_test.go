package calsum

import (
	"fmt"
	"testing"
	"time"
)

func TestCalsum(t *testing.T) {
	in := make(chan string)
	out := make(chan string)
	err := revoke(in, out)
	fmt.Println(err)
	//in <- "10"
	//fmt.Println(<-out)
	//
	startPipe := time.Now()
	for i := 1; i < 50000; i++ {
		in <- fmt.Sprintf("%d", i)
		<-out
	}
	costPipe := time.Now().Sub(startPipe)

	fmt.Printf("cycle: 500000, pipe: %s, pipe/cycle: %s\n", costPipe, costPipe/time.Duration(500000))

	cycles := []int{50000, 100000, 500000, 1000000}
	counts := []int{10, 50, 100, 500, 1000, 5000, 10000}
	for _, count := range counts {
		for _, cycle := range cycles {
			startCgo := time.Now()
			for i := 0; i < cycle; i++ {
				CgocalSum(count)
			}
			costCgo := time.Now().Sub(startCgo)

			startGo := time.Now()
			for i := 0; i < cycle; i++ {
				GoCalSum(count)
			}
			costGo := time.Now().Sub(startGo)

			fmt.Printf("count: %d, cycle: %d, cgo: %s, go: %s, cgo/cycle: %s, go/cycle: %s, cgo/go: %.4f \n",
				count, cycle, costCgo, costGo,
				costCgo/time.Duration(cycle), costGo/time.Duration(cycle),
				float64(costCgo)/float64(costGo))
		}
	}
}
