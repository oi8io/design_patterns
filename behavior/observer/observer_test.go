package behavior

import "testing"

func TestObserver(t *testing.T) {
	i := newItem("iphone")
	c1 := newCustomer("steven chow")
	c2 := newCustomer("andy lau")
	c3 := newCustomer("Jacky Cheung")
	c4 := newCustomer("Yun Fat chow")
	i.register(c1)
	i.register(c2)
	i.notifyAll()



	j := newItem("android")
	j.register(c3)
	j.register(c2)
	j.register(c4)
	j.notifyAll()


	j.setName("android 8")
	j.deRegister(c4)
	j.deRegister(c1)
	j.notifyAll()

	i.setName("iphone 14")
	i.register(c4)
	i.notifyAll()
}
