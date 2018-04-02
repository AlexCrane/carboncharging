package batterystate

import (
	"github.com/AlexCrane/carboncharging/actuator"
)

type DummyBatteryState struct {
}

func NewDummyBatteryState() BatteryState {
	return &DummyBatteryState{}
}

func (dbs *DummyBatteryState) IsConnected(actuator actuator.Actuator) bool {
	return true
}

func (dbs *DummyBatteryState) GetCurrentCharge() Milliwatts {
	return Milliwatts(100)
}

func (dbs *DummyBatteryState) GetFullCapacity() Milliwatts {
	return Milliwatts(200)
}

func (dbs *DummyBatteryState) GetBatteryDraw() Milliwatts {
	return Milliwatts(1)
}

func (dbs *DummyBatteryState) GetComputerDraw() Milliwatts {
	return Milliwatts(1)
}

func (dbs *DummyBatteryState) IsBatteryFull() bool {
	return false
}
