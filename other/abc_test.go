package other

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

func sequenceABC() {
	chanA, chanB, chanC := make(chan struct{}, 0), make(chan struct{}, 0), make(chan struct{}, 0)
	n := 3

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			<-chanA
			fmt.Println("A")
			chanB <- struct{}{}
		}
		<-chanA
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			<-chanB
			fmt.Println("B")
			chanC <- struct{}{}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			<-chanC
			fmt.Println("C")
			chanA <- struct{}{}
		}
	}()

	chanA <- struct{}{}
	wg.Wait()
}

func TestSequenceABC(t *testing.T) {
	sequenceABC()
}

func TestSizeOf(t *testing.T) {
	a, b, c, d := 1, 2, 3, 4
	arr := []*int{&a, &b, &c, &d, &a, &b, &c, &d}
	arr2 := [...]*int{&a, &b, &c, &d}
	fmt.Println(unsafe.Sizeof(arr))
	fmt.Println(unsafe.Sizeof(arr2))
}
