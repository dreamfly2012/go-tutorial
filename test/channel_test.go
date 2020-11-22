package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {

	channel := make(chan bool)
	go func() {
		time.Sleep(1 * time.Second)
		<-channel
	}()
	start := time.Now()
	channel <- true
	fmt.Printf("send took %v\n", time.Since(start))
}

func TestDeadLock(t *testing.T) {

	c := make(chan bool, 1)
	c <- true
	<-c
}

func TestDoWork(t *testing.T) {
	c := make(chan int)

	for i := 0; i < 4; i++ {
		go doWork(c)
	}
	for {
		v := <-c
		fmt.Println(v)
	}

}

func doWork(c chan int) {
	for {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		c <- rand.Int()
	}
}

func TestChannel2(t *testing.T) {
	c := make(chan bool)
	for i := 0; i < 5; i++ {
		go func(x int) {
			sendRPC(x)
			c <- true
		}(i)
	}
}

func sendRPC(x int) {
	println(x)
}
