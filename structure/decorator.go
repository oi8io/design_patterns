package structure

/**
装饰器模式
 */
type pizza interface {
	getPrice() int
}

type veggeMania struct {

}

func (v *veggeMania) getPrice() int {
	return 15
}

type tomatoTopping struct {
	pizza pizza
}

func (t *tomatoTopping) getPrice() int {
	return t.pizza.getPrice()+9
}

type cheeseTopping struct {
	pizza pizza
}

func (c *cheeseTopping) getPrice() int {
	return c.pizza.getPrice()+10
}



