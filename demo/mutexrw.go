package demo

import (
	"fmt"
	"sync"
)

var (
	data  = make(map[string]string)
	mutex sync.RWMutex
)








func read(key string, wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.RLock()
	value := data[key]
	fmt.Println("Read:", key, "=", value)

	mutex.RUnlock()
}

func write(key, value string, wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock()
	fmt.Println("Writing:", key, "=", value)
	data[key] = value

	mutex.Unlock()
}

func mai() {
	var wg sync.WaitGroup

	wg.Add(2)
	go write("name", "Alice", &wg)
	go write("city", "Chennai", &wg)

	
	wg.Add(3)
	go read("name", &wg)
	go read("city", &wg)
	go read("name", &wg)

	wg.Wait()
}
