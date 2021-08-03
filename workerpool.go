package workerpool

import (
	"sync"
)

type workerPool struct {
	mutex sync.Mutex
	ch    chan bool
	errs  []error
}

func (w *workerPool) Submit(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				w.mutex.Lock()
				w.errs = append(w.errs, err.(error))
				w.mutex.Unlock()
			}
			<-w.ch
		}()
		w.ch <- true
		fn()
	}()
}

func (w *workerPool) Wait() []error {
	for {
		if len(w.ch) == 0 {
			close(w.ch)
			return w.errs
		}
	}
}

func NewWorkerPool(maxWorkerNumber int) *workerPool {
	return &workerPool{
		ch:   make(chan bool, maxWorkerNumber),
		errs: make([]error, 0),
	}
}
