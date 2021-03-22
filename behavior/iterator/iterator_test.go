package behavior

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	ui := new(userIterator)
	for i := 0; i < 10; i++ {
		ui.users = append(ui.users, newUser(fmt.Sprintf("user %d", i), 10+i))
	}
	for ui.hasNext() {
		fmt.Println(ui.getNext())
	}
}
