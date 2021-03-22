package behavior

import (
	"fmt"
)

/**
中介者莫斯
*/
type train interface {
	arrive()
	depart()
	permitArrival()
}

type passengerTrain struct {
	name     string
	mediator mediator
}

func newPassengerTrain(name string, mediator mediator) *passengerTrain {
	return &passengerTrain{name: name, mediator: mediator}
}

func (t *passengerTrain) arrive() {
	if !t.mediator.canArrive(t) {
		fmt.Printf("[%s]: Arrival blocked, waiting\n", t.name)
	} else {
		fmt.Printf("[%s]: Arrived \n", t.name)
	}
}

func (t *passengerTrain) depart() {
	fmt.Printf("[%s]: Leaving\n", t.name)
	t.mediator.notifyAboutDeparture()
}

func (t *passengerTrain) permitArrival() {
	fmt.Printf("[%s]: Arrival permitted, arriving\n ", t.name)
	t.arrive()
}

type freightTrain struct {
	name     string
	mediator mediator
}

func newFreightTrain(name string, mediator mediator) *freightTrain {
	return &freightTrain{name: name, mediator: mediator}
}

func (t *freightTrain) arrive() {
	if !t.mediator.canArrive(t) {
		fmt.Printf("[%s]: Arrival blocked, waiting\n", t.name)
	} else {
		fmt.Printf("[%s]: Arrived \n", t.name)
	}
}

func (t *freightTrain) depart() {
	fmt.Printf("[%s]: Leaving\n", t.name)
	t.mediator.notifyAboutDeparture()
}

func (t *freightTrain) permitArrival() {
	fmt.Printf("[%s]: Arrival permitted, arriving\n ", t.name)
	t.arrive()
}

type mediator interface {
	canArrive(t train) bool
	notifyAboutDeparture()
}

type stationManager struct {
	isPlatformFree bool
	trainQueue     []train
}

func newStationManager() *stationManager {
	return &stationManager{isPlatformFree: true}
}

func (m *stationManager) canArrive(t train) bool {
	if m.isPlatformFree { // 空的直接开进来
		m.isPlatformFree = false
		return true
	}
	m.trainQueue = append(m.trainQueue, t)
	return false
}

func (m *stationManager) notifyAboutDeparture() {
	if !m.isPlatformFree {
		m.isPlatformFree = true
	}
	if len(m.trainQueue) > 0 {
		firstTrainInQueue := m.trainQueue[0]
		m.trainQueue = m.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}
