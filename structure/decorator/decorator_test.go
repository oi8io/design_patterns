package structure

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	pizz := &veggeMania{}
	pizzWithChese := &cheeseTopping{pizza: pizz}
	pizzaWithTomatoTopping := &tomatoTopping{pizza: pizzWithChese}
	price := pizzaWithTomatoTopping.getPrice()
	fmt.Println(price)
}
