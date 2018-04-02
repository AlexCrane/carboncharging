package chargescheduler

import (
	"time"

	"github.com/AlexCrane/carboncharging/actuator"
)

type ChargeScheduler struct {
	schedule     *ChargeSchedule
	actuator     actuator.Actuator
	scheduleDone chan bool

	interrupt chan bool
}

func NewChargeScheduler(schedule *ChargeSchedule, actuator actuator.Actuator) *ChargeScheduler {
	cs := &ChargeScheduler{
		schedule:     schedule,
		actuator:     actuator,
		scheduleDone: make(chan bool, 1),
		interrupt:    make(chan bool, 0),
	}

	cs.start()

	return cs
}

func (cs *ChargeScheduler) start() {
	go func() {
		done := true // this will be set to false if the current set of inflexion points is interrupted by an update
		inflexionPoints := cs.schedule.GetInflexionPoints()
		for _, inflexionPoint := range inflexionPoints {
			select {
			case <-cs.interrupt:
				done = false
				break
			case <-time.After(inflexionPoint.At.Sub(time.Now())):
				if inflexionPoint.ChargeOn {
					cs.actuator.StartCharging()
				} else {
					cs.actuator.StopCharging()
				}
			}
		}

		if done {
			select {
			case <-cs.interrupt:
				done = false
			case <-time.After(cs.schedule.schedulePeriods[len(cs.schedule.schedulePeriods)-1].to.Sub(time.Now())):
				done = true
			}

			if done {
				cs.scheduleDone <- true
			}
		}
	}()
}

func (cs *ChargeScheduler) UpdateSchedule(schedule *ChargeSchedule) {
	cs.schedule = schedule
	cs.interrupt <- true
	cs.start()
}

func (cs *ChargeScheduler) WaitForDone() {
	<-cs.scheduleDone
}
