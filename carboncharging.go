package main

import (
	"log"

	"github.com/AlexCrane/carboncharging/actuator"
	"github.com/AlexCrane/carboncharging/batterystate"
	"github.com/AlexCrane/carboncharging/chargescheduler"
	"github.com/AlexCrane/carboncharging/chargestrategy"
)

func main() {
	actuator := actuator.NewDummyActuator()
	batteryState := batterystate.NewDiStatusBatteryState(actuator)

	log.Println("IsConnected", batteryState.IsConnected(actuator))
	log.Println("GetCurrentCharge", batteryState.GetCurrentCharge())
	log.Println("GetFullCapacity", batteryState.GetFullCapacity())
	log.Println("GetBatteryDraw", batteryState.GetBatteryDraw())
	log.Println("GetComputerDraw", batteryState.GetComputerDraw())
	log.Println("IsBatteryFull", batteryState.IsBatteryFull())

	if batteryState.IsConnected(actuator) {
		strategy := chargestrategy.NewDummyStrategy()

		chargeScheduler := chargescheduler.NewChargeScheduler(strategy.GetCurrentChargeSchedule(), actuator)

		go func() {
			updates := strategy.GetChargeScheduleChan()

			for {
				schedule, more := <-updates
				if more {
					log.Println("Updated schedule!")
					chargeScheduler.UpdateSchedule(schedule)
				}
			}
		}()

		chargeScheduler.WaitForDone()
		log.Println("Schedule done!")
	}
}
