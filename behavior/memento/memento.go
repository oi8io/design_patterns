package behavior

type memento struct {
	state string
}

func newMemento(state string) *memento {
	return &memento{state: state}
}

func (m *memento) getSaveState() string {
	return m.state
}

type eMemento interface {
	createMemento() (m *memento)
	restoreMemento(m *memento)
	setState(state string)
	getState() (state string)
}

type originator struct {
	state string
}

func newOriginator(state string) *originator {
	return &originator{state: state}
}

func (o *originator) createMemento() (m *memento) {
	return newMemento(o.state)
}

func (o *originator) restoreMemento(m *memento) {
	o.state = m.getSaveState()
}

func (o *originator) setState(state string) {
	o.state = state
}

func (o *originator) getState() (state string) {
	return o.state
}

type caretaker struct {
	mementoArray []*memento
}

func newCaretaker() *caretaker {
	return &caretaker{mementoArray: []*memento{}}
}

func (t *caretaker) addMemento(m *memento) {
	t.mementoArray = append(t.mementoArray, m)
}

func (t *caretaker) getMemento(i int) (m *memento) {
	if len(t.mementoArray) > i {
		return t.mementoArray[i]
	}
	return nil
}
