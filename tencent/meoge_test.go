package tencent

import (
	"fmt"
	"sync"
	"testing"
)

// 两个协程交替打印字符

// N 10
// A B A B 重复

func printAAndB(n int) {
	chanA := make(chan struct{})
	chanB := make(chan struct{})
	numA, numB := 0, 0
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for {
			<-chanA
			fmt.Println("A")
			numA++
			chanB <- struct{}{}
			if numA == n {
				wg.Done()
				return
			}
		}
	}()

	go func() {
		for {
			<-chanB
			fmt.Println("B")
			numB++
			if numB == n {
				wg.Done()
				return
			}
			chanA <- struct{}{}
		}
	}()

	chanA <- struct{}{}
	wg.Wait()
}

func Test_printAB(t *testing.T) {
	printAAndB(10)
}


// 1M * 8Byte = 8MB?  num top100 