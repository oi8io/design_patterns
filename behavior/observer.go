package behavior

import "fmt"

type subject interface {
	register(server ObServer)
	deRegister(server ObServer)
	notifyAll()
}

type item struct {
	observerMap map[string]ObServer
	name        string
	inSock      bool
}

func (i *item) register(server ObServer) {
	i.observerMap[server.getID()] = server
}

func (i *item) deRegister(server ObServer) {
	delete(i.observerMap, server.getID())
}

func (i *item) notifyAll() {
	for _, server := range i.observerMap {
		server.update(i.name)
	}
}

func (i *item) setName(name string) {
	i.name = name
}

func newItem(name string) *item {
	return &item{name: name, observerMap: make(map[string]ObServer)}
}

type ObServer interface {
	update(string)
	getID() string
}

type customer struct {
	id string
}

func newCustomer(id string) *customer {
	return &customer{id: id}
}

func (c *customer) update(s string) {
	fmt.Printf("Sending email to customer [%18s] for item [%s]\n", c.id, s)
}

func (c *customer) getID() string {
	return c.id
}
