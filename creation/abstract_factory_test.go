package creation

import (
	"fmt"
	"testing"
)

// https://refactoringguru.cn/design-patterns/abstract-factory/go/example
// https://golangbyexample.com/abstract-factory-design-pattern-go/
func TestGetSportsFactoryNike(t *testing.T) {
	fa := GetSportsFactory(FactoryNike)
	shirt := fa.MakeShirt()
	shoe := fa.MakeShoe()
	shoe.setLogo("Adidas")  // overwise
	shirt.setLogo("Adidas") // overwise
	fmt.Println(shirt.toString())
	fmt.Println(shoe.toString())
}

func TestGetSportsFactoryAdadis(t *testing.T) {
	fa := GetSportsFactory(FactoryAdidas)
	shirt := fa.MakeShirt()
	shoe := fa.MakeShoe()
	shoe.setLogo("Nike")  //blocking
	shirt.setLogo("Nike") // blocking
	fmt.Println(shirt.toString())
	fmt.Println(shoe.toString())
}
