package structure

import "fmt"

type Composite interface {
	process(space string)
}

type Component struct {
	name    string
	element []Composite
}

func (c *Component) process(space string) {
	fmt.Printf("%sComponent [%s] process\n",space, c.name)
	for _, i2 := range c.element {
		i2.process(space + space)
	}
}

func (c *Component) Add(com Composite) {
	c.element = append(c.element, com)
}

type Leaf struct {
	name string
}

func (l *Leaf) process(space string) {
	fmt.Printf("%sLeaf [%s] process\n", space, l.name)
}
