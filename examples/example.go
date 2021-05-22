package main

import (
	"fmt"
	"time"

	"github.com/yoavgur/changroup"
)

func main() {
	// Create changroup
	var cg changroup.ChanGroup

	// Add and send to goroutine
	cg.Add(1)
	// Call goroutine, function should finish after 10 seconds
	go doSomething(&cg)

	// Wait on waitgroup with timeout to assure behaviour won't block forever
	select {
	case <-cg.WaitCh():
		fmt.Println("Finished waiting for goroutine")

	case <-time.After(time.Second * 3):
		fmt.Println("Waiting timed out")
	}
}

func doSomething(cg *changroup.ChanGroup) {
	defer cg.Done()

	for i := 0; i < 10; i++ {
		fmt.Println("Doing")
		time.Sleep(time.Second * 1)
	}
}
