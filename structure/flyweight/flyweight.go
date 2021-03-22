package structure

import (
	"errors"
)

/**
享元模式
 */
type terroristDress struct { //恐怖服装
	color string
}

func (t *terroristDress) getColor() string {
	return t.color
}

func newTerroristDress() *terroristDress {
	return &terroristDress{color: "red"}
}

type counterTerroristDress struct { //反恐服装
	color string
}

func (c *counterTerroristDress) getColor() string {
	return c.color
}

func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{color: "green"}
}

type dress interface {
	getColor() string
}

type dressFactory struct {
	dressMap map[string]dress
}

const (
	TerroristDressType        = "t"
	CounterTerroristDressType = "ct"
)

var dressFactorySingleInstance = &dressFactory{dressMap: make(map[string]dress)}

func getDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}
func (d *dressFactory) getDressByType(dressType string) (dress, error) {
	if dt, ok := d.dressMap[dressType]; ok {
		return dt, nil
	}
	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}
	return nil, errors.New("params error")
}

type player struct {
	dress      dress  // 服装
	playerType string // 可为 T 或 CT
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &player{playerType: playerType, dress: dress}
}

func (p *player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

type game struct {
	Terrorists []*player
	CounterTerrorists []*player
}

func newGame() *game {
	return &game{}
}

func (g *game) addTerrorist(dressType string )  {
	g.Terrorists = append(g.Terrorists, newPlayer(dressType,dressType))
}
func (g *game) addCounterTerrorist(dressType string )  {
	g.CounterTerrorists = append(g.CounterTerrorists, newPlayer(dressType,dressType))
}

