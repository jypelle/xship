package main

import (
	"time"
	"../../internal"
)

var iterationStartTime time.Time

func main() {

	//	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// Setup devices
	internal.Setup()

	// Start main loop
	var duration time.Duration
	for {
		iterationStartTime = time.Now()

		internal.Update()

		internal.Draw()

		duration = 32 * time.Millisecond - time.Now().Sub(iterationStartTime)
		if duration>0 {
			time.Sleep(duration)
		}
	}
}
