package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, exists := sm.data[key]
	return val, exists
}

func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

func (sm *SafeMap) Len() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return len(sm.data)
}

func testSafeMapRW() {
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			key := fmt.Sprintf("key_%d", idx)
			safeMap.Set(key, idx)
		}(i)
	}

	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			key := fmt.Sprintf("key_%d", idx%50)
			if val, exists := safeMap.Get(key); exists {
				_ = val
			}
		}(i)
	}

	wg.Wait()
}

func testConcurrentWrites() {
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(2)

		go func(writerID int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("writer_%d_key_%d", writerID, j)
				safeMap.Set(key, j)
				time.Sleep(time.Microsecond * 10)
			}
		}(i)

		go func(readerID int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("writer_%d_key_%d", readerID%10, j)
				safeMap.Get(key)
				time.Sleep(time.Microsecond * 5)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final map size: %d\n", safeMap.Len())
}

func main() {
	fmt.Println("\nSafeMapRW test with sync.RWMutex")
	testSafeMapRW()

	fmt.Println("\nConcurrency test")
	testConcurrentWrites()
}
