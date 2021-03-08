package creation

import (
	"fmt"
	"sync"
)

var once sync.Once
var mu sync.Mutex

type single struct {
}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("create single")
			singleInstance = new(single)
		})
	}
	return singleInstance
}
func GetInstanceV2() *single {
	if singleInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		if singleInstance == nil {
			fmt.Println("create single v2")
			singleInstance = new(single)
		}
	}
	return singleInstance
}
