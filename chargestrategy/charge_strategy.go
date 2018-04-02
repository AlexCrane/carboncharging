package chargestrategy

import "github.com/AlexCrane/carboncharging/chargescheduler"

type ChargeStrategy interface {
	// GetCurrentChargeSchedule gets the current charge schedule
	GetCurrentChargeSchedule() *chargescheduler.ChargeSchedule

	// GetChargeScheduleChan gets a channel that can push an updated ChargeSchedule allowing for a dynamic strategy
	GetChargeScheduleChan() <-chan *chargescheduler.ChargeSchedule
}
