package changroup

import "sync"

type ChanGroup struct {
	sync.WaitGroup
	wait   chan struct{}
	doOnce sync.Once
}

func (cg *ChanGroup) WaitCh() <-chan struct{} {
	cg.doOnce.Do(func() {
		cg.wait = make(chan struct{})

		go func() {
			cg.Wait()
			close(cg.wait)
		}()
	})

	return cg.wait
}
