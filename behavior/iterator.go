package behavior

type collection interface {
	create() iterator
}

type iterator interface {
	hasNext() bool
	getNext() interface{}
}
type user struct {
	name string
	age  int
}

func newUser(name string, age int) *user {
	return &user{name: name, age: age}
}

type userIterator struct {
	cursor int
	users  []*user
}

func (u *userIterator) hasNext() bool {
	if u.cursor < len(u.users)-1 {
		return true
	}
	return false
}

func (u *userIterator) getNext() interface{} {
	if u2 := u.users[u.cursor]; u2 != nil {
		u.cursor++
		return u2
	}

	return nil
}
