package wait

import (
	"sync"
)

type Wait struct {
	control sync.WaitGroup
	count   int
	block   bool
	mutex   sync.RWMutex
}

func New() *Wait {

	return &Wait{
		control: sync.WaitGroup{},
		count:   0,
		block:   false,
		mutex:   sync.RWMutex{},
	}
}

func (w *Wait) Add() bool {

	w.mutex.Lock()
	defer w.mutex.Unlock()

	if w.block {
		return false
	}
	w.control.Add(1)
	w.count += 1
	return true
}

func (w *Wait) Done() bool {

	w.mutex.Lock()
	defer w.mutex.Unlock()
	if w.count == 0 {
		return false
	}
	w.control.Done()
	w.count -= 1
	return true
}

func (w *Wait) Count() int {

	w.mutex.RLock()
	defer w.mutex.RUnlock()
	return w.count
}

func (w *Wait) Wait() {

	w.control.Wait()
}

func (w *Wait) IsBlock() bool {

	w.mutex.Lock()
	defer w.mutex.Unlock()
	return w.block
}

func (w *Wait) Block() {

	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.block = true
}
