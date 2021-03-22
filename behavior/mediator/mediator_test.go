package behavior

import "testing"

func TestMediator(t *testing.T) {
	mediator := newStationManager()
	passengerTrain1 := newPassengerTrain("passengerTrain1", mediator)
	passengerTrain2 := newPassengerTrain("passengerTrain2", mediator)
	passengerTrain3 := newPassengerTrain("passengerTrain3", mediator)
	passengerTrain4 := newPassengerTrain("passengerTrain4", mediator)
	freightTrain1 := newFreightTrain("freightTrain1", mediator)
	freightTrain2 := newFreightTrain("freightTrain2", mediator)
	freightTrain3 := newFreightTrain("freightTrain3", mediator)
	freightTrain4 := newFreightTrain("freightTrain4", mediator)
	passengerTrain1.arrive()
	passengerTrain2.arrive()
	freightTrain1.arrive()
	freightTrain3.permitArrival()
	passengerTrain3.permitArrival()
	freightTrain2.arrive()
	freightTrain4.arrive()
	passengerTrain1.depart()
	passengerTrain2.depart()
	freightTrain1.depart()
	freightTrain3.depart()
	passengerTrain3.depart()
	freightTrain2.depart()
	passengerTrain4.permitArrival()
	freightTrain4.depart()
	passengerTrain4.depart()

}
