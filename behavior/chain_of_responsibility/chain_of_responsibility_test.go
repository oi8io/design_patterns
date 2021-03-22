package behavior

import (
	"testing"
)

func TestChain(t *testing.T) {

	reception := newReception()
	cashier := newCashier()
	bloodRoutine := newBloodRoutine()
	electrocardiogram := newElectrocardiogram()
	chest := newChest()
	colposcopy := newColposcopy()
	urine := newUrine()
	reception.setNext(cashier)
	cashier.setNext(bloodRoutine)
	bloodRoutine.setNext(electrocardiogram)
	electrocardiogram.setNext(chest)
	chest.setNext(colposcopy)
	colposcopy.setNext(urine)
	p := &patient{number: 1232}
	reception.execute(p)
}
