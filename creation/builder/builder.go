package creation

import "fmt"

/**
将对象的构建和实现进行分离

构建汽车步骤比较 『稳定』
具体实施细节中各有不同，属于变化
对稳定的进行抽象
对变化进行实现
*/
type Builder interface {
	createVehicle() // 创建车辆
	addWheel()
	addEngine()
	addDoors()
	getVehicle() string //获取车辆
}

const (
	CarTypeCar  = "Car"
	CarTypeBike = "Bike"
)

type Vehicle struct {
	name   string
	wheel  int //轮子数量
	engine string
	doors  int
}

type CarBuilder struct {
	Vehicle
}

func (c *CarBuilder) createVehicle() {
	c.name = CarTypeCar
}

func (c *CarBuilder) addWheel() {
	c.wheel = 4 // 4个滚滚
}

func (c *CarBuilder) addEngine() {
	c.engine = "8缸2.0T排量发动机"
}

func (c *CarBuilder) addDoors() {
	c.doors = 4
}

func (c *CarBuilder) getVehicle() string {
	return fmt.Sprintf("%+v", c)
}

type BikeBuilder struct {
	Vehicle
}

func (b *BikeBuilder) createVehicle() {
	b.name = CarTypeBike
}

func (b *BikeBuilder) addWheel() {
	b.wheel = 2
}

func (b *BikeBuilder) addEngine() {
	b.engine = "防滑脚踏板"
}

func (b *BikeBuilder) addDoors() {
	b.doors = 0
}

func (b *BikeBuilder) getVehicle() string {
	return fmt.Sprintf("%+v", b)
}

type VehicleDirector struct {
}

func (dr *VehicleDirector) Build(builder Builder) {
	builder.createVehicle()
	builder.addDoors()
	builder.addWheel()
	builder.addEngine()
	result := builder.getVehicle()
	fmt.Println(result)
}
