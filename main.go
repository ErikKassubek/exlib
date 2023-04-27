package exlib

import (
	"fmt"
	"sync"
)

// ============= Mutexe =============

func MutexInValue(mu sync.Mutex) {
	mu.Lock()
	fmt.Println("MutexInValue")
	mu.Unlock()
}

func MutexInReference(mu *sync.Mutex) {
	mu.Lock()
	fmt.Println("MutexInReference")
	mu.Unlock()
}

func MutexInOutValue(mu sync.Mutex) sync.Mutex {
	fmt.Println("MutexInOutValue")
	return mu

}

func MutexInOutReference(mu *sync.Mutex) *sync.Mutex {
	fmt.Println("MutexInOutReference")
	return mu

}

// ============= RW-Mutexe =============
func RWMutexInValue(mu sync.RWMutex) {
	mu.Lock()
	fmt.Println("RWMutexInValue")
	mu.Unlock()
}

func RWMutexInReference(mu *sync.RWMutex) {
	mu.Lock()
	fmt.Println("RWMutexInReference")
	mu.Unlock()
}

func RWMutexInOutValue(mu sync.RWMutex) sync.RWMutex {
	fmt.Println("RWMutexInOutValue")
	return mu

}

func RWMutexInOutReference(mu *sync.RWMutex) *sync.RWMutex {
	fmt.Println("RWMutexInOutReference")
	return mu

}

// ============= channel =============

func ChannelInValue(ch chan int) {
	fmt.Println("ChannelInValue")
}

func ChannelInReference(ch *chan int) {
	fmt.Println("ChannelInReference")
}

func ChannelInOutValue(ch chan int) chan int {
	fmt.Println("ChannelInOutValue")
	return ch
}

func ChannelInOutReference(ch *chan int) *chan int {
	fmt.Println("ChannelInOutReference")
	return ch
}
