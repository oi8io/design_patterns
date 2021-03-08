package creation

import (
	"fmt"
	"testing"
)

func TestSingleTon(t *testing.T) {
	for i := 0; i < 30; i++ {
		go GetInstanceV2()
	}
	go func() {
		for i := 0; i < 30; i++ {
			go GetInstance()
		}
	}()
	fmt.Scanln()
}
