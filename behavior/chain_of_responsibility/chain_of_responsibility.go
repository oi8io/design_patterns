package behavior

import (
	"fmt"
	"time"
)

/**
体检的时候，医师体检完一个项目，就让你下一个项目去多少号
*/
type department interface {
	execute(p *patient)
	setNext(next department)
}

type patient struct {
	number            int64
	reception         bool // 心电图
	cashier           bool // 心电图
	bloodRoutine      bool // 血常规
	electrocardiogram bool // 心电图
	chest             bool // 胸部正位
	colposcopy        bool // 阴道镜检查
	urine             bool // 尿常规
}

type urine struct {
	name string
	next department
}

func newUrine() *urine {
	return &urine{name: "urine"}
}

func (d *urine) execute(p *patient) {
	p.urine = true
	Log(p.number, d.name)
	if d.next != nil {
		d.next.execute(p)
	}
}
func Log(num int64, name string) {
	fmt.Printf("[%s] patient [%d] do  [%s] done \n", time.Now().Format(time.RFC3339), num, name)
}

func (d *urine) setNext(next department) {
	d.next = next
}

type reception struct {
	name string
	next department
}

func newReception() *reception {
	return &reception{name: "reception"}
}

func (d *reception) execute(p *patient) {
	p.reception = true
	Log(p.number, d.name)
	if d.next != nil {
		d.next.execute(p)
	}
}

func (d *reception) setNext(next department) {
	d.next = next
}

type cashier struct {
	name string
	next department
}

func newCashier() *cashier {
	return &cashier{name: "Cashier"}
}

func (d *cashier) execute(p *patient) {
	p.cashier = true
	Log(p.number, d.name)
	if d.next != nil {
		d.next.execute(p)
	}
}

func (d *cashier) setNext(next department) {
	d.next = next
}

type colposcopy struct {
	name string
	next department
}

func newColposcopy() *colposcopy {
	return &colposcopy{name: "Colposcopy"}
}

func (d *colposcopy) execute(p *patient) {
	p.colposcopy = true
	Log(p.number, d.name)
	if d.next != nil {
		d.next.execute(p)
	}
}

func (d *colposcopy) setNext(next department) {
	d.next = next
}

type bloodRoutine struct {
	name string
	next department
}

func newBloodRoutine() *bloodRoutine {
	return &bloodRoutine{name: "bloodRoutine"}
}

func (d *bloodRoutine) execute(p *patient) {
	p.bloodRoutine = true
	Log(p.number, d.name)
	if d.next != nil {
		d.next.execute(p)
	}
}

func (d *bloodRoutine) setNext(next department) {
	d.next = next
}

type electrocardiogram struct {
	name string
	next department
}

func newElectrocardiogram() *electrocardiogram {
	return &electrocardiogram{name: "electrocardiogram"}
}

func (d *electrocardiogram) execute(p *patient) {
	p.electrocardiogram = true
	Log(p.number, d.name)
	if d.next != nil {
		d.next.execute(p)
	}
}

func (d *electrocardiogram) setNext(next department) {
	d.next = next
}

type chest struct {
	name string
	next department
}

func newChest() *chest {
	return &chest{name: "chest"}
}

func (d *chest) execute(p *patient) {
	p.electrocardiogram = true
	Log(p.number, d.name)
	if d.next != nil {
		d.next.execute(p)
	}
}

func (d *chest) setNext(next department) {
	d.next = next
}
