package batterystate

import (
	"fmt"

	"github.com/AlexCrane/carboncharging/actuator"
)

type Milliwatts int

func (m Milliwatts) String() string {
	return fmt.Sprintf("%dmW", int(m))
}

type BatteryState interface {
	IsConnected(actuator actuator.Actuator) bool

	// Suspect we need lots of stuff here:

	GetCurrentCharge() Milliwatts
	GetFullCapacity() Milliwatts

	GetBatteryDraw() Milliwatts
	GetComputerDraw() Milliwatts

	IsBatteryFull() bool
}
