package behavior

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	caretaker:=newCaretaker()
	originator :=newOriginator("A")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("B")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("C")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

    originator.restoreMemento(caretaker.getMemento(2))
	fmt.Printf("Originator Current State: %s\n", originator.getState())
    originator.restoreMemento(caretaker.getMemento(0))
	fmt.Printf("Originator Current State: %s\n", originator.getState())
    originator.restoreMemento(caretaker.getMemento(1))
	fmt.Printf("Originator Current State: %s\n", originator.getState())

}