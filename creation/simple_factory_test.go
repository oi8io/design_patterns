package creation

import (
	"fmt"
	"testing"
)

func TestNewSimpleFactoryStorage(t *testing.T) {
	st := NewSimpleFactoryStorage(StorageTypeMySQL)
	st.Create(1)
	fmt.Println(st.Read())
	st.Create("+_+_+__+)")
	fmt.Println(st.Read())
	st.Delete("asdf")
	fmt.Println(st.Read())
}
