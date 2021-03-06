// Peripheral: ADC_Periph  Analog to Digital Converter.
// Instances:
//  ADC1  mmap.ADC1_BASE
// Registers:
//  0x00 32  ISR    Interrupt and status register.
//  0x04 32  IER    Interrupt enable register.
//  0x08 32  CR     Control register.
//  0x0C 32  CFGR1  Configuration register 1.
//  0x10 32  CFGR2  Configuration register 2.
//  0x14 32  SMPR   Sampling time register.
//  0x20 32  TR     Analog watchdog 1 threshold register.
//  0x28 32  CHSELR Group regular sequencer register.
//  0x40 32  DR     Group regular data register.
// Import:
//  stm32/o/f030x6/mmap
package adc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	ADRDY ISR = 0x01 << 0 //+ ADC ready flag.
	EOSMP ISR = 0x01 << 1 //+ ADC group regular end of sampling flag.
	EOC   ISR = 0x01 << 2 //+ ADC group regular end of unitary conversion flag.
	EOS   ISR = 0x01 << 3 //+ ADC group regular end of sequence conversions flag.
	OVR   ISR = 0x01 << 4 //+ ADC group regular overrun flag.
	AWD1  ISR = 0x01 << 7 //+ ADC analog watchdog 1 flag.
)

const (
	ADRDYn = 0
	EOSMPn = 1
	EOCn   = 2
	EOSn   = 3
	OVRn   = 4
	AWD1n  = 7
)

const (
	ADRDYIEIE IER = 0x01 << 0 //+ ADC ready interrupt.
	EOSMPIEIE IER = 0x01 << 1 //+ ADC group regular end of sampling interrupt.
	EOCIEIE   IER = 0x01 << 2 //+ ADC group regular end of unitary conversion interrupt.
	EOSIEIE   IER = 0x01 << 3 //+ ADC group regular end of sequence conversions interrupt.
	OVRIEIE   IER = 0x01 << 4 //+ ADC group regular overrun interrupt.
	AWD1IEIE  IER = 0x01 << 7 //+ ADC analog watchdog 1 interrupt.
)

const (
	ADRDYIEIEn = 0
	EOSMPIEIEn = 1
	EOCIEIEn   = 2
	EOSIEIEn   = 3
	OVRIEIEn   = 4
	AWD1IEIEn  = 7
)

const (
	ADEN    CR = 0x01 << 0  //+ ADC enable.
	ADDIS   CR = 0x01 << 1  //+ ADC disable.
	ADSTART CR = 0x01 << 2  //+ ADC group regular conversion start.
	ADSTP   CR = 0x01 << 4  //+ ADC group regular conversion stop.
	ADCAL   CR = 0x01 << 31 //+ ADC calibration.
)

const (
	ADENn    = 0
	ADDISn   = 1
	ADSTARTn = 2
	ADSTPn   = 4
	ADCALn   = 31
)

const (
	DMAEN   CFGR1 = 0x01 << 0  //+ ADC DMA transfer enable.
	DMACFG  CFGR1 = 0x01 << 1  //+ ADC DMA transfer configuration.
	SCANDIR CFGR1 = 0x01 << 2  //+ ADC group regular sequencer scan direction.
	RES     CFGR1 = 0x03 << 3  //+ ADC data resolution.
	ALIGN   CFGR1 = 0x01 << 5  //+ ADC data alignement.
	EXTSEL  CFGR1 = 0x07 << 6  //+ ADC group regular external trigger source.
	EXTEN   CFGR1 = 0x03 << 10 //+ ADC group regular external trigger polarity.
	OVRMOD  CFGR1 = 0x01 << 12 //+ ADC group regular overrun configuration.
	CONT    CFGR1 = 0x01 << 13 //+ ADC group regular continuous conversion mode.
	WAIT    CFGR1 = 0x01 << 14 //+ ADC low power auto wait.
	AUTOFF  CFGR1 = 0x01 << 15 //+ ADC low power auto power off.
	DISCEN  CFGR1 = 0x01 << 16 //+ ADC group regular sequencer discontinuous mode.
	AWD1SGL CFGR1 = 0x01 << 22 //+ ADC analog watchdog 1 monitoring a single channel or all channels.
	AWD1EN  CFGR1 = 0x01 << 23 //+ ADC analog watchdog 1 enable on scope ADC group regular.
	AWD1CH  CFGR1 = 0x1F << 26 //+ ADC analog watchdog 1 monitored channel selection.
)

const (
	DMAENn   = 0
	DMACFGn  = 1
	SCANDIRn = 2
	RESn     = 3
	ALIGNn   = 5
	EXTSELn  = 6
	EXTENn   = 10
	OVRMODn  = 12
	CONTn    = 13
	WAITn    = 14
	AUTOFFn  = 15
	DISCENn  = 16
	AWD1SGLn = 22
	AWD1ENn  = 23
	AWD1CHn  = 26
)

const (
	CKMODE CFGR2 = 0x03 << 30 //+ ADC clock source and prescaler (prescaler only for clock source synchronous).
)

const (
	CKMODEn = 30
)

const (
	SMP SMPR = 0x07 << 0 //+ ADC group of channels sampling time 2.
)

const (
	SMPn = 0
)

const (
	LT1 TR = 0xFFF << 0  //+ ADC analog watchdog 1 threshold low.
	HT1 TR = 0xFFF << 16 //+ ADC Analog watchdog 1 threshold high.
)

const (
	LT1n = 0
	HT1n = 16
)

const (
	CHSEL   CHSELR = 0x7FFFF << 0 //+ ADC group regular sequencer channels, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL18 CHSELR = 0x40000 << 0 //  ADC group regular sequencer channel 18, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL17 CHSELR = 0x20000 << 0 //  ADC group regular sequencer channel 17, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL16 CHSELR = 0x10000 << 0 //  ADC group regular sequencer channel 16, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL15 CHSELR = 0x8000 << 0  //  ADC group regular sequencer channel 15, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL14 CHSELR = 0x4000 << 0  //  ADC group regular sequencer channel 14, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL13 CHSELR = 0x2000 << 0  //  ADC group regular sequencer channel 13, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL12 CHSELR = 0x1000 << 0  //  ADC group regular sequencer channel 12, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL11 CHSELR = 0x800 << 0   //  ADC group regular sequencer channel 11, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL10 CHSELR = 0x400 << 0   //  ADC group regular sequencer channel 10, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL9  CHSELR = 0x200 << 0   //  ADC group regular sequencer channel 9, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL8  CHSELR = 0x100 << 0   //  ADC group regular sequencer channel 8, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL7  CHSELR = 0x80 << 0    //  ADC group regular sequencer channel 7, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL6  CHSELR = 0x40 << 0    //  ADC group regular sequencer channel 6, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL5  CHSELR = 0x20 << 0    //  ADC group regular sequencer channel 5, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL4  CHSELR = 0x10 << 0    //  ADC group regular sequencer channel 4, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL3  CHSELR = 0x08 << 0    //  ADC group regular sequencer channel 3, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL2  CHSELR = 0x04 << 0    //  ADC group regular sequencer channel 2, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL1  CHSELR = 0x02 << 0    //  ADC group regular sequencer channel 1, available when ADC_CFGR1_CHSELRMOD is reset.
	CHSEL0  CHSELR = 0x01 << 0    //  ADC group regular sequencer channel 0, available when ADC_CFGR1_CHSELRMOD is reset.
)

const (
	CHSELn = 0
)

const (
	DATA DR = 0xFFFF << 0 //+ ADC group regular conversion data.
)

const (
	DATAn = 0
)
