package test

import (
	"sync"
	"testing"
	"time"
)

var done bool

var mu sync.Mutex

func TestMutex(t *testing.T) {
	time.Sleep(1 * time.Second)
	println("started")
	go periodic()
	time.Sleep(5 * time.Second)
	mu.Lock()
	done = true
	mu.Unlock()
	println("canceld")
	time.Sleep(3 * time.Second)
}

func TestMutex2(t *testing.T) {
	var count = 0
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			count = count + 1
			mu.Unlock()
		}()
	}
	time.Sleep(1 * time.Second)

	println(count)
	t.Log("finished")
}

func periodic() {
	for {
		println("tick")
		time.Sleep(1 * time.Second)
		//mu.Lock()
		if done {
			//mu.Unlock()
			return
		}
		//mu.Unlock()
	}
}
