package actuator

import "log"

// DummyActuator is an implementation of Actuator which prints instructions to plug in and unplug charger to the command line
type DummyActuator struct {
	isCharging bool
}

// NewDummyActuator returns a new DummyActuator instance
func NewDummyActuator() Actuator {
	return &DummyActuator{
		isCharging: false,
	}
}

// StartCharging tells the user to plug in charger
func (da *DummyActuator) StartCharging() error {
	log.Println("Plug in charger!")
	da.isCharging = true
	return nil
}

// StopCharging tells the user to plug in charger
func (da *DummyActuator) StopCharging() error {
	log.Println("Unplug charger!")
	da.isCharging = false
	return nil
}

func (da *DummyActuator) IsCharging() bool {
	return da.isCharging
}
