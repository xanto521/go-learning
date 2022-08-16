package main

//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//// SafeCounter is a safe wrapper for a Counter.
//type SafeCounter struct {
//	v   map[string]int
//	mux sync.Mutex
//}
//
//// Inc increments the counter for the given key.
//func (c *SafeCounter) Inc(key string) {
//	c.mux.Lock()
//	// Lock so only one goroutine at a time can access the map c.v.
//	c.v[key]++
//	c.mux.Unlock()
//}
//
//func (c *SafeCounter) Value(key string) int {
//	c.mux.Lock()
//	// Lock so only one goroutine at a time can access the map c.v.
//	defer c.mux.Unlock()
//	return c.v[key]
//}
//
//func main() {
//	c := SafeCounter{v: make(map[string]int)}
//	for i := 0; i < 1000; i++ {
//		go c.Inc("somekey")
//	}
//	// if not sleep, the program will exit before the goroutines finish.
//	time.Sleep(1 * time.Millisecond)
//	fmt.Println(c.Value("somekey"))
//}
