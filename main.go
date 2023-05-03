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

// ============= channel =============

func ChannelInValueSend(ch chan int) {
	fmt.Println("Send 1")
	ch <- 1
	fmt.Println("Send 2")
}

func ChannelInReferenceSend(ch *chan int) {
	*ch <- 1
}

func ChannelInValueReceive(ch chan int) {
	<-ch
}

func ChannelInReferenceReceive(ch *chan int) {
	<-*ch
}
