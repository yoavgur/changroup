package changroup

import "sync"

type ChanGroup struct {
	sync.WaitGroup
	wait   chan struct{}
	doOnce sync.Once
}

// WaitCh returns a channel which will be closed when the waitgroup counter is zero
// When called it creates a single goroutine which waits for the waitgroup and closes the channel.
// This function can be called multiple times, but once the waitgroup has finished waiting
// re-calling this function will always return a closed channel. Thus you cannot reuse a finished waitgroup.
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
