// +build f10x_hd

// Package mmap provides base memory adresses for all peripherals.
package mmap

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	FLASH_BASE     uintptr = 0x08000000 // FLASH base address in the alias region
	SRAM_BASE      uintptr = 0x20000000 // SRAM base address in the alias region
	PERIPH_BASE    uintptr = 0x40000000 // Peripheral base address in the alias region
	SRAM_BB_BASE   uintptr = 0x22000000 // SRAM base address in the bit-band region
	PERIPH_BB_BASE uintptr = 0x42000000 // Peripheral base address in the bit-band region
	FSMC_R_BASE    uintptr = 0xA0000000 // FSMC registers base address
)

// Peripheral memory map
const (
	APB1PERIPH_BASE    uintptr = PERIPH_BASE
	APB2PERIPH_BASE    uintptr = PERIPH_BASE + 0x10000
	AHBPERIPH_BASE     uintptr = PERIPH_BASE + 0x20000
	TIM2_BASE          uintptr = APB1PERIPH_BASE + 0x0000
	TIM3_BASE          uintptr = APB1PERIPH_BASE + 0x0400
	TIM4_BASE          uintptr = APB1PERIPH_BASE + 0x0800
	TIM5_BASE          uintptr = APB1PERIPH_BASE + 0x0C00
	TIM6_BASE          uintptr = APB1PERIPH_BASE + 0x1000
	TIM7_BASE          uintptr = APB1PERIPH_BASE + 0x1400
	TIM12_BASE         uintptr = APB1PERIPH_BASE + 0x1800
	TIM13_BASE         uintptr = APB1PERIPH_BASE + 0x1C00
	TIM14_BASE         uintptr = APB1PERIPH_BASE + 0x2000
	RTC_BASE           uintptr = APB1PERIPH_BASE + 0x2800
	WWDG_BASE          uintptr = APB1PERIPH_BASE + 0x2C00
	IWDG_BASE          uintptr = APB1PERIPH_BASE + 0x3000
	SPI2_BASE          uintptr = APB1PERIPH_BASE + 0x3800
	SPI3_BASE          uintptr = APB1PERIPH_BASE + 0x3C00
	USART2_BASE        uintptr = APB1PERIPH_BASE + 0x4400
	USART3_BASE        uintptr = APB1PERIPH_BASE + 0x4800
	UART4_BASE         uintptr = APB1PERIPH_BASE + 0x4C00
	UART5_BASE         uintptr = APB1PERIPH_BASE + 0x5000
	I2C1_BASE          uintptr = APB1PERIPH_BASE + 0x5400
	I2C2_BASE          uintptr = APB1PERIPH_BASE + 0x5800
	CAN1_BASE          uintptr = APB1PERIPH_BASE + 0x6400
	CAN2_BASE          uintptr = APB1PERIPH_BASE + 0x6800
	BKP_BASE           uintptr = APB1PERIPH_BASE + 0x6C00
	PWR_BASE           uintptr = APB1PERIPH_BASE + 0x7000
	DAC_BASE           uintptr = APB1PERIPH_BASE + 0x7400
	CEC_BASE           uintptr = APB1PERIPH_BASE + 0x7800
	AFIO_BASE          uintptr = APB2PERIPH_BASE + 0x0000
	EXTI_BASE          uintptr = APB2PERIPH_BASE + 0x0400
	GPIOA_BASE         uintptr = APB2PERIPH_BASE + 0x0800
	GPIOB_BASE         uintptr = APB2PERIPH_BASE + 0x0C00
	GPIOC_BASE         uintptr = APB2PERIPH_BASE + 0x1000
	GPIOD_BASE         uintptr = APB2PERIPH_BASE + 0x1400
	GPIOE_BASE         uintptr = APB2PERIPH_BASE + 0x1800
	GPIOF_BASE         uintptr = APB2PERIPH_BASE + 0x1C00
	GPIOG_BASE         uintptr = APB2PERIPH_BASE + 0x2000
	ADC1_BASE          uintptr = APB2PERIPH_BASE + 0x2400
	ADC2_BASE          uintptr = APB2PERIPH_BASE + 0x2800
	TIM1_BASE          uintptr = APB2PERIPH_BASE + 0x2C00
	SPI1_BASE          uintptr = APB2PERIPH_BASE + 0x3000
	TIM8_BASE          uintptr = APB2PERIPH_BASE + 0x3400
	USART1_BASE        uintptr = APB2PERIPH_BASE + 0x3800
	ADC3_BASE          uintptr = APB2PERIPH_BASE + 0x3C00
	TIM15_BASE         uintptr = APB2PERIPH_BASE + 0x4000
	TIM16_BASE         uintptr = APB2PERIPH_BASE + 0x4400
	TIM17_BASE         uintptr = APB2PERIPH_BASE + 0x4800
	TIM9_BASE          uintptr = APB2PERIPH_BASE + 0x4C00
	TIM10_BASE         uintptr = APB2PERIPH_BASE + 0x5000
	TIM11_BASE         uintptr = APB2PERIPH_BASE + 0x5400
	SDIO_BASE          uintptr = PERIPH_BASE + 0x18000
	DMA1_BASE          uintptr = AHBPERIPH_BASE + 0x0000
	DMA1_Channel1_BASE uintptr = AHBPERIPH_BASE + 0x0008
	DMA1_Channel2_BASE uintptr = AHBPERIPH_BASE + 0x001C
	DMA1_Channel3_BASE uintptr = AHBPERIPH_BASE + 0x0030
	DMA1_Channel4_BASE uintptr = AHBPERIPH_BASE + 0x0044
	DMA1_Channel5_BASE uintptr = AHBPERIPH_BASE + 0x0058
	DMA1_Channel6_BASE uintptr = AHBPERIPH_BASE + 0x006C
	DMA1_Channel7_BASE uintptr = AHBPERIPH_BASE + 0x0080
	DMA2_BASE          uintptr = AHBPERIPH_BASE + 0x0400
	DMA2_Channel1_BASE uintptr = AHBPERIPH_BASE + 0x0408
	DMA2_Channel2_BASE uintptr = AHBPERIPH_BASE + 0x041C
	DMA2_Channel3_BASE uintptr = AHBPERIPH_BASE + 0x0430
	DMA2_Channel4_BASE uintptr = AHBPERIPH_BASE + 0x0444
	DMA2_Channel5_BASE uintptr = AHBPERIPH_BASE + 0x0458
	RCC_BASE           uintptr = AHBPERIPH_BASE + 0x1000
	CRC_BASE           uintptr = AHBPERIPH_BASE + 0x3000
	FLASH_R_BASE       uintptr = AHBPERIPH_BASE + 0x2000 // Flash registers base address
	OB_BASE            uintptr = 0x1FFFF800              // Flash Option Bytes base address
	ETH_BASE           uintptr = AHBPERIPH_BASE + 0x8000
	ETH_MAC_BASE       uintptr = ETH_BASE
	ETH_MMC_BASE       uintptr = ETH_BASE + 0x0100
	ETH_PTP_BASE       uintptr = ETH_BASE + 0x0700
	ETH_DMA_BASE       uintptr = ETH_BASE + 0x1000
	FSMC_Bank1_R_BASE  uintptr = FSMC_R_BASE + 0x0000 // FSMC Bank1 registers base address
	FSMC_Bank1E_R_BASE uintptr = FSMC_R_BASE + 0x0104 // FSMC Bank1E registers base address
	FSMC_Bank2_R_BASE  uintptr = FSMC_R_BASE + 0x0060 // FSMC Bank2 registers base address
	FSMC_Bank3_R_BASE  uintptr = FSMC_R_BASE + 0x0080 // FSMC Bank3 registers base address
	FSMC_Bank4_R_BASE  uintptr = FSMC_R_BASE + 0x00A0 // FSMC Bank4 registers base address
	DBGMCU_BASE        uintptr = 0xE0042000           // Debug MCU registers base address
)
