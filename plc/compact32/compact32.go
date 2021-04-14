package compact32

import (
	"mylab/cpagent/config"
	"mylab/cpagent/plc"
	"time"

	"github.com/goburrow/modbus"
	logger "github.com/sirupsen/logrus"
)

const (
	HOLDING_STAGE = "hold"
	CYCLE_STAGE   = "cycle"
)



type Compact32 struct {
	ExitCh  chan error
	WsMsgCh chan string
	wsErrch chan error
	Driver  plc.Compact32Driver
}



func NewCompact32Driver(wsMsgCh chan string, wsErrch chan error, exit chan error, test bool) plc.Driver {
	/* Modbus RTU/ASCII */
	handler := modbus.NewRTUClientHandler(config.ReadEnvString("MODBUS_TTY"))
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "E"
	handler.StopBits = 1
	handler.SlaveId = 5 // THis is hard-coded as the PLC RS485 is configured as SlaveID-5
	handler.Timeout = 500 * time.Millisecond

	handler.Connect()
	driver := Compact32ModbusDriver{}
	driver.Client = modbus.NewClient(handler)

	C32 := Compact32{}
	C32.Driver = &driver
	C32.ExitCh = exit
	C32.WsMsgCh = wsMsgCh

	// Start the Heartbeat
	// TODO: Uncomment this after RT-PCR m/c is ready
	// go C32.HeartBeat()

	// Specifically for testing!
	if test {
		p := plc.Stage{
			Holding: []plc.Step{
				plc.Step{65.3, 3.2, 1},
				plc.Step{85.3, 3.1, 1},
				plc.Step{95, 2.8, 1},
			},
			Cycle: []plc.Step{
				plc.Step{55, 3, 1},
				plc.Step{65, 3, 1},
				plc.Step{75, 3, 1},
				plc.Step{85, 3, 1},
				plc.Step{95, 3, 1},
			},
			CycleCount: 3,
		}

		C32.Stop()
		C32.ConfigureRun(p)
		C32.Start()

		cycle := uint16(0)
		for {
			scan, err := C32.Monitor(cycle)
			if err != nil {
				logger.WithField("error", err).Error("Error in Monitoring")
				break
			}
			// Log everytime to scan changes!
			if scan.Cycle != cycle {
				cycle = scan.Cycle
				logger.WithField("scan", scan).Info("Monitoring..")
			}
			logger.WithField("temperature", scan.Temp).Info("Monitoring..")

			if cycle == p.CycleCount {
				// Cycling done .. issue stop!
				break
			}
			time.Sleep(200 * time.Millisecond)
		}
		C32.Stop()
	}

	return &C32 // plc Driver
}

// Compact32 Driver for Deck A and B
// TODO: Use test to Configure the Deck Operations
func NewCompact32DeckDriverA(wsMsgCh chan string, wsErrch chan error, exit chan error, test bool) (plc.Common, *modbus.RTUClientHandler) {
	/* Modbus RTU/ASCII */
	handler := modbus.NewRTUClientHandler(config.ReadEnvString("MODBUS_TTY"))
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "E"
	handler.StopBits = 1
	handler.SlaveId = byte(1)
	handler.Timeout = 50 * time.Millisecond

	handler.Connect()
	driver := Compact32ModbusDriver{}

	driver.Client = modbus.NewClient(handler)

	C32 := plc.Compact32Deck{}
	C32.DeckDriver = &driver
	C32.ExitCh = exit
	C32.WsMsgCh = wsMsgCh
	C32.WsErrCh = wsErrch

	// C32.Name = "A"
	plc.SetDeckName(&C32, "A")


	return &C32, handler // plc Driver
}

func NewCompact32DeckDriverB(wsMsgCh chan string, exit chan error, test bool, handler *modbus.RTUClientHandler) plc.Common {
	/* Modbus RTU/ASCII */
	handler2 := modbus.NewRTUClientHandler(config.ReadEnvString("MODBUS_TTY"))
	handler2.BaudRate = 9600
	handler2.DataBits = 8
	handler2.Parity = "E"
	handler2.StopBits = 1
	handler2.SlaveId = byte(2)
	handler2.Timeout = 50 * time.Millisecond

	handler2.Connect()
	driver := Compact32ModbusDriver{}
	driver.Client = modbus.NewClient2(handler2, handler)

	C32 := plc.Compact32Deck{}
	C32.DeckDriver = &driver
	C32.ExitCh = exit
	C32.WsMsgCh = wsMsgCh

	plc.SetDeckName(&C32, "B")

	// C32.Name = "B"

	return &C32 // plc Driver
}
