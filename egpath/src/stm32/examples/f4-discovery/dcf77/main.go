package main

import (
	"fmt"
	"rtos"
	"time"

	"dcf77"
	"stm32/f4/exti"
	"stm32/f4/gpio"
	"stm32/f4/irq"
	"stm32/f4/periph"
	"stm32/f4/setup"
)

func init() {
	setup.Performance168(8)

	initLEDs()

	// Initialize DCF77 input pin.

	periph.APB2ClockEnable(periph.SysCfg)
	periph.APB2Reset(periph.SysCfg)
	periph.AHB1ClockEnable(periph.GPIOC)
	periph.AHB1Reset(periph.GPIOC)

	gpio.C.SetMode(1, gpio.In)
	exti.L1.Connect(gpio.C)
	exti.L1.RiseTrigEnable()
	exti.L1.FallTrigEnable()
	exti.L1.IntEnable()
	rtos.IRQ(irq.Ext1).Enable()

	periph.APB2ClockDisable(periph.SysCfg)
}

var d = dcf77.NewDecoder()

func edgeISR() {
	t := time.Now()
	exti.L1.ClearPending()
	blink(Blue, -100)
	d.Edge(t, gpio.C.InPin(1) != 0)
}

var ISRs = [...]func(){
	irq.Ext1: edgeISR,
} //c:__attribute__((section(".ISRs")))


func main() {
	for {
		p := d.Pulse()
		now := time.Now().UnixNano()
		if p.Err() != nil {
			fmt.Printf("now=%d %v\n", now, p.Err())
			blink(Red, 25)
		} else {
			stamp := p.Stamp.UnixNano()
			fmt.Printf("now=%d stamp=%d dcf=%s\n", now, stamp, p.Date)
			blink(Green, 25)
		}
	}
}
