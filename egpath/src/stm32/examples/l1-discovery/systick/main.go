// This example shows how to manually setup interrupt table when you don't use
// runtime initialisation (MaxTasks == 0) and how to write purely interrupt
// driven application.
package main

import (
	"arch/cortexm"
	"arch/cortexm/scb"
	"arch/cortexm/systick"

	"stm32/l1/gpio"
	"stm32/l1/irq"
	"stm32/l1/periph"
	"stm32/l1/setup"
)

var LED = gpio.B

const (
	Blue  = 6
	Green = 7
)

func defaultHandler() {
	for {
	}
}

var (
	cnt   int
	ledup = true
)

func isr() {
	if ledup {
		LED.SetPin(Blue)
	} else {
		LED.ClearPin(Blue)
	}
	ledup = !ledup
}

var IRQs = [...]func(){
	irq.Tim2: isr,
} //c:__attribute__((section(".ISRs")))

func main() {
	setup.Performance(0)

	periph.AHBClockEnable(periph.GPIOB)
	periph.AHBReset(periph.GPIOB)

	LED.SetMode(Blue, gpio.Out)
	LED.SetMode(Green, gpio.Out)

	tenms := systick.TENMS.Load()
	tenms *= 10 // stm32l1 returns value for 1 ms not for 10ms.
	systick.RELOAD.Store(tenms * 100)
	systick.CURRENT.Clear()
	(systick.ENABLE | systick.TICKINT).Set()

	// Sleep forever.
	scb.SLEEPONEXIT.Set()
	cortexm.DSB() // not necessary on Cortex-M0,M3,M4
	cortexm.WFI()

	// Execution should never reach there so the green LED
	// should never light up.
	LED.SetPin(Green)
}
