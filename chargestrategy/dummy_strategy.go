package chargestrategy

import (
	"time"

	"github.com/AlexCrane/carboncharging/chargescheduler"
)

type DummyStrategy struct {
	currentSchedule    *chargescheduler.ChargeSchedule
	scheduleUpdateChan chan *chargescheduler.ChargeSchedule
}

func NewDummyStrategy() ChargeStrategy {
	chargeSchedule := chargescheduler.NewChargeSchedule()
	chargeSchedule.AddSchedulePeriod(time.Now(), time.Now().Add(time.Second*5), true)
	chargeSchedule.AddSchedulePeriod(time.Now().Add(time.Second*5), time.Now().Add(time.Second*15), false)
	chargeSchedule.AddSchedulePeriod(time.Now().Add(time.Second*15), time.Now().Add(time.Second*30), true)

	ds := &DummyStrategy{
		currentSchedule:    chargeSchedule,
		scheduleUpdateChan: make(chan *chargescheduler.ChargeSchedule, 0),
	}

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 10)

			chargeSchedule := chargescheduler.NewChargeSchedule()
			chargeSchedule.AddSchedulePeriod(time.Now(), time.Now().Add(time.Second*5), true)
			chargeSchedule.AddSchedulePeriod(time.Now().Add(time.Second*5), time.Now().Add(time.Second*15), false)
			chargeSchedule.AddSchedulePeriod(time.Now().Add(time.Second*15), time.Now().Add(time.Second*30), true)

			ds.currentSchedule = chargeSchedule
			ds.scheduleUpdateChan <- ds.currentSchedule
		}

		close(ds.scheduleUpdateChan)
	}()

	return ds
}

// GetCurrentChargeSchedule gets the current charge schedule
func (ds *DummyStrategy) GetCurrentChargeSchedule() *chargescheduler.ChargeSchedule {
	return ds.currentSchedule
}

// GetChargeScheduleChan gets a channel that can push an updated ChargeSchedule allowing for a dynamic strategy
func (ds *DummyStrategy) GetChargeScheduleChan() <-chan *chargescheduler.ChargeSchedule {
	return ds.scheduleUpdateChan
}
