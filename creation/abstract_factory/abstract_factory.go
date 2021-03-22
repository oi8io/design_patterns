package creation

import "fmt"

// https://refactoringguru.cn/design-patterns/abstract-factory/go/example
// https://golangbyexample.com/abstract-factory-design-pattern-go/

type AbstractShoe interface {
	getLogo() string
	getSize() int
	setSize(size int)
	setLogo(logo string)
	toString() string
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) getSize() int {
	return s.size
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}
func (s *Shoe) toString() string {
	return fmt.Sprintf("this is your shoe: %+v", s)
}

type ShirtSize struct {
	size   string //袖长
	sleeve int    //袖长
	chest  int    //袖长
	body   int    //袖长
}

const (
	TSizeS  = "S"
	TSizeM  = "M"
	TSizeL  = "L"
	TSizeXL = "XL"
)

func NewTSize(sizeName string) ShirtSize {
	switch sizeName {
	case TSizeS:
		return ShirtSize{size: sizeName, sleeve: 25, chest: 40, body: 105}
	case TSizeM:
		return ShirtSize{size: sizeName, sleeve: 30, chest: 42, body: 108}
	case TSizeL:
		return ShirtSize{size: sizeName, sleeve: 35, chest: 45, body: 112}
	case TSizeXL:
		return ShirtSize{size: sizeName, sleeve: 35, chest: 45, body: 112}
	}
	return ShirtSize{size: sizeName, sleeve: 30, chest: 42, body: 108}
}

type AbstractShirt interface {
	getLogo() string
	getSize() ShirtSize
	setSize(sizeName string)
	setLogo(logo string)
	toString() string
}

type Shirt struct {
	logo string
	size ShirtSize
}

func (s *Shirt) setLogo(logo string) {
	panic("implement me")
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) getSize() ShirtSize {
	return s.size
}

func (s *Shirt) setSize(sizeName string) {
	s.size = NewTSize(sizeName)
}
func (s *Shirt) toString() string {
	return fmt.Sprintf("this is your thirt: %+v", s)
}

type AdidasShoe struct {
	Shoe
}

type AdidasShirt struct {
	Shirt
}

type NikeShoe struct {
	Shoe
}
type NikeShirt struct {
	Shirt
}

func (a *AdidasShoe) getLogo() string {
	return "adidas"
}
func (a *AdidasShoe) setLogo(logo string) {
}
func (a *AdidasShirt) getLogo() string {
	return "adidas"
}
func (a *AdidasShirt) setLogo(logo string) {

}

func (a *NikeShirt) getLogo() string {
	return "nike"
}
func (a *NikeShirt) setLogo(logo string) {
}
func (a *NikeShoe) getLogo() string {
	return "nike"
}
func (a *NikeShoe) setLogo(logo string) {
}

type AbstractSportFactory interface {
	MakeShoe() AbstractShoe
	MakeShirt() AbstractShirt
}

type defaultFactory struct {
}

func (d defaultFactory) MakeShoe() AbstractShoe {
	panic("implement me")
}

func (d defaultFactory) MakeShirt() AbstractShirt {
	panic("implement me")
}

type FactoryBrandAdidas struct {
}

func (b *FactoryBrandAdidas) MakeShoe() AbstractShoe {
	return &AdidasShoe{Shoe{
		logo: "adidas",
		size: 10,
	}}
}

func (b *FactoryBrandAdidas) MakeShirt() AbstractShirt {
	return &AdidasShirt{Shirt{
		logo: "adidas", size: NewTSize(TSizeM),
	}}
}

type FactoryBrandNike struct {
}

func (b *FactoryBrandNike) MakeShoe() AbstractShoe {
	return &NikeShoe{Shoe{
		logo: "adidas", size: 10,
	}}
}

func (b *FactoryBrandNike) MakeShirt() AbstractShirt {
	return &NikeShirt{Shirt{
		logo: "nike", size: NewTSize(TSizeM),
	}}
}

const (
	FactoryNike   = "nike"
	FactoryAdidas = "adidas"
)

func GetSportsFactory(name string) AbstractSportFactory {
	switch name {
	case FactoryNike:
		return &FactoryBrandNike{}
	case FactoryAdidas:
		return &FactoryBrandAdidas{}
	}
	return &defaultFactory{}
}
