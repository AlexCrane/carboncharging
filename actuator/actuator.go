package actuator

// Actuator is an interface for module that start and stop charging
type Actuator interface {
	StartCharging() error
	StopCharging() error

	IsCharging() bool
}
