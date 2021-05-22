package changroup

import "sync"

type ChanGroup struct {
	wg *sync.WaitGroup
}

func (cg *ChanGroup) Add(delta int) {
	cg.wg.Add(delta)
}

func (cg *ChanGroup) Done() {
	cg.wg.Done()
}

func (cg *ChanGroup) Wait() <-chan struct{} {
	wait := make(chan struct{})

	go func() {
		cg.wg.Wait()
		close(wait)
	}()

	return wait
}
