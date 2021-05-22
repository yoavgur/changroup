package changroup

import "sync"

type ChanGroup struct {
	sync.WaitGroup
}

func (cg *ChanGroup) WaitCh() <-chan struct{} {
	wait := make(chan struct{})

	go func() {
		cg.Wait()
		close(wait)
	}()

	return wait
}
