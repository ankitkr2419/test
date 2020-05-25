package compact32

import (
	"github.com/goburrow/modbus"
	"mylab/cpagent/plc"
	"time"
)

type Compact32 struct {
	Client modbus.Client
}

var compact32 Compact32 = Compact32{}

func NewCompact32Driver() plc.Driver {
	/* Modbus RTU/ASCII */
	handler := modbus.NewRTUClientHandler("/dev/tty.usbserial-A50285BI") // TODO: Read from config file or flag!
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "E"
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Timeout = 2 * time.Second

	handler.Connect()
	compact32.Client = modbus.NewClient(handler)

	return &compact32
}

func (d *Compact32) HeartBeat() error {
	_, err := compact32.Client.WriteSingleRegister(MODBUS["D"][100], uint16(2)) //write

	return err
}

func (d *Compact32) PreRun(plc.Stage) error {
	return nil
}

// Monitor periodically. If Status=CYCLE_COMPLETE, the Scan will be populated
func (d *Compact32) Monitor() (scan plc.Scan, status plc.Status) {
	return
}
