package workerpool

import (
	"sync"
)

type workerPool struct {
	mutex sync.Mutex
	ch    chan bool
	errs  []error
}

func (s *workerPool) Submit(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				s.mutex.Lock()
				s.errs = append(s.errs, err.(error))
				s.mutex.Unlock()
			}
			<-s.ch
		}()
		s.ch <- true
		fn()
	}()
}

func (s *workerPool) Wait() []error {
	for {
		if len(s.ch) == 0 {
			close(s.ch)
			return s.errs
		}
	}
}

func NewWorkerPool(maxWorkerNumber int) *workerPool {
	return &workerPool{
		ch:   make(chan bool, maxWorkerNumber),
		errs: make([]error, 0),
	}
}
