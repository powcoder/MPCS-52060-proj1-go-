https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
// Package lock provides an implementation of a read-write lock
// that uses condition variables and mutexes.
package lock

import "sync"


type RWLock struct {
	
	RN int
	WN int	
	
	Mutex *sync.Mutex
	Cond *sync.Cond
}

func NewRWLock() *RWLock {
	mutex := new(sync.Mutex)
	cond := sync.NewCond(mutex)
	
	return &RWLock {
		Mutex: mutex,
		Cond: cond,
		RN: 0,
		WN: 0,
	}
	
}

func (this *RWLock) RLock() {
	
	this.Mutex.Lock()
	
	for this.WN > 0 {
		this.Cond.Wait()
	}
	
	this.RN++	
	
	this.Mutex.Unlock()
	
}

func (this *RWLock) RUnLock() {
	
	this.Mutex.Lock()
	
	this.RN--
	
	this.Cond.Broadcast()
	
	
	this.Mutex.Unlock()
}

func (this *RWLock) Lock() {
	
	this.Mutex.Lock()	
	
	for this.WN != 0 || this.RN != 0 {
		this.Cond.Wait()
	}
	
	this.WN++
	
	this.Mutex.Unlock()
	
}

func (this *RWLock) UnLock() {
	this.Mutex.Lock()
	
	this.WN--	
	
	this.Cond.Broadcast()
	
	this.Mutex.Unlock()
	
}