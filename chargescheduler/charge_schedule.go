package chargescheduler

import (
	"time"
)

type SchedulePoint struct {
	At       time.Time
	ChargeOn bool
}

type schedulePeriod struct {
	from     time.Time
	to       time.Time
	chargeOn bool
}

type ChargeSchedule struct {
	schedulePeriods []*schedulePeriod
}

func NewChargeSchedule() *ChargeSchedule {
	return &ChargeSchedule{
		schedulePeriods: make([]*schedulePeriod, 0),
	}
}

func (cs *ChargeSchedule) AddSchedulePeriod(from time.Time, to time.Time, chargeOn bool) {
	cs.schedulePeriods = append(cs.schedulePeriods, &schedulePeriod{
		from:     from,
		to:       to,
		chargeOn: chargeOn,
	})
}

func (cs *ChargeSchedule) GetInflexionPoints() []*SchedulePoint {
	if len(cs.schedulePeriods) == 0 {
		return make([]*SchedulePoint, 0)
	}

	// Strictly fewer inflexion points than schedule periods - two consecutive schedule periods may have
	// the same 'chargeOn' so wouldn't be an inflexion point
	inflexion := make([]*SchedulePoint, 0, len(cs.schedulePeriods))
	inflexion = append(inflexion, &SchedulePoint{At: cs.schedulePeriods[0].from,
		ChargeOn: cs.schedulePeriods[0].chargeOn})

	lastChargeOn := cs.schedulePeriods[0].chargeOn

	for _, schedulePeriod := range cs.schedulePeriods[1:] {
		if schedulePeriod.chargeOn != lastChargeOn {
			inflexion = append(inflexion, &SchedulePoint{At: schedulePeriod.from,
				ChargeOn: schedulePeriod.chargeOn})
			lastChargeOn = schedulePeriod.chargeOn
		}
	}

	return inflexion
}
