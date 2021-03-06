package main

import (
	"fmt"
	"sync"
)

var (
	// mu   = sync.RWMutex{}
	initOnce = sync.Once{}
	data     map[string][]string
)

func expensiveInit() {
	data = make(map[string][]string)
	data["drink"] = []string{"coke", "sprite", "bubble tea", "coffee", "tea"}
	data["fruit"] = []string{"orange", "watermelon", "rose apple"}
	data["food"] = []string{"Pad Thai", "Basil fry egg"}

	fmt.Println("Expensive initialized")
}

func lookUp(catagory string) []string {
	// mu.RLock()
	// if data != nil {
	// 	mu.RUnlock()
	// 	return data[catagory]
	// }

	// mu.RUnlock()

	// mu.Lock()
	// if data == nil {

	// 	expensiveInit()
	// }
	// mu.Unlock()
	// return data[catagory]

	initOnce.Do(expensiveInit)
	return data[catagory]

}
func main() {

	x := "drink"
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lookUp(x)
		}()

	}
	wg.Wait()
	fmt.Println("Done")
}
