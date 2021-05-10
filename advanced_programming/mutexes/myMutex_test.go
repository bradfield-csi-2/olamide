package myMutex

import (
	//"fmt"
	"sync"
	"testing"
)

func testFunc(c int, m *MyMutex, wg *sync.WaitGroup) {
	defer wg.Done()
	m.lock()
	c++
	//fmt.Println("The current count is: ", c)
	m.unlock()
}

func testFunc1(c int, m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Lock()
	c++
	//fmt.Println("The current count is: ", c)
	m.Unlock()
}

func TestMyMutex(t *testing.T) {
	//description := "Test the mutex library"
	mu := MyMutex{}
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go testFunc(i, &mu, &wg)
	}

	wg.Wait()
}

func TestMutex(t *testing.T) {
	//description := "Test the mutex library"
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go testFunc1(i, &mu, &wg)
	}

	wg.Wait()
}

func BenchmarkMyMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mu := MyMutex{}
		var wg sync.WaitGroup
		for i := 0; i < 1; i++ {
			wg.Add(1)
			go testFunc(i, &mu, &wg)
		}
	}
}

func BenchmarkMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mu := MyMutex{}
		var wg sync.WaitGroup
		for i := 0; i < 1; i++ {
			wg.Add(1)
			go testFunc(i, &mu, &wg)
		}
	}
}
