package batterystate

import (
	"log"
	"time"

	"github.com/AlexCrane/carboncharging/actuator"
	"github.com/distatus/battery"
)

type DiStatusBatteryState struct {
	actuator  actuator.Actuator
	batteries []*battery.Battery
}

func NewDiStatusBatteryState(actuator actuator.Actuator) BatteryState {
	batteries, err := battery.GetAll()
	if err != nil {
		log.Fatalf("Can't create DiStateBatteryState; %v", err)
	}

	return &DiStatusBatteryState{
		actuator:  actuator,
		batteries: batteries,
	}
}

func (dbs *DiStatusBatteryState) IsConnected(actuator actuator.Actuator) bool {
	// A rather heroic assumption...
	return actuator == dbs.actuator
}

func (dbs *DiStatusBatteryState) GetCurrentCharge() Milliwatts {
	var currentCharge Milliwatts

	for _, batt := range dbs.batteries {
		currentCharge += Milliwatts(batt.Current)
	}

	return currentCharge
}

func (dbs *DiStatusBatteryState) GetFullCapacity() Milliwatts {
	var currentCapacity Milliwatts

	for _, batt := range dbs.batteries {
		currentCapacity += Milliwatts(batt.Full)
	}

	return currentCapacity
}

func (dbs *DiStatusBatteryState) GetBatteryDraw() Milliwatts {
	var currentDraw Milliwatts

	for _, batt := range dbs.batteries {
		if batt.State == battery.Charging {
			currentDraw += Milliwatts(batt.ChargeRate)
		}
	}

	return currentDraw
}

func (dbs *DiStatusBatteryState) GetComputerDraw() Milliwatts {
	var currentDraw Milliwatts

	toggleCharging := dbs.actuator.IsCharging()
	if toggleCharging {
		dbs.actuator.StopCharging()
		time.Sleep(100 * time.Millisecond)
	}

	for _, batt := range dbs.batteries {
		if batt.State == battery.Discharging {
			currentDraw += Milliwatts(batt.ChargeRate)
		}
	}

	if toggleCharging {
		dbs.actuator.StartCharging()
	}

	return currentDraw
}

func (dbs *DiStatusBatteryState) IsBatteryFull() bool {
	full := true

	for _, batt := range dbs.batteries {
		if batt.State != battery.Full {
			full = false
			break
		}
	}

	return full
}
