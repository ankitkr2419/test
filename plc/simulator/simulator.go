package simulator

import (
	"fmt"
	"math/rand"
	"mylab/cpagent/plc"
	"time"

	logger "github.com/sirupsen/logrus"
)

type plcRegistors struct {
	d holdingRegistors
	m coilRegistors
}
type holdingRegistors struct {
	heartbeat      uint16    // heartbeat register (W)
	currentTemp    uint16    // Current temperature (R)
	currentCycle   uint16    // Current cycle (R)
	idealLidTemp   uint16    // Ideal Lid temperature (W)
	currentLidTemp uint16    // Current Lid temperature (R)
	emissions      [6]uint16 // Well Emission data 96x6 registers (R)
	stage          plc.Stage // contains cycle data
	errCode        uint16    // error code (R)
}
type coilRegistors struct {
	EmmissionFlag  uint16 // Well Emmission register data  ON: PLC write & OFF: Read (RW)
	cycleCompleted uint16 // Cycle completed (R)
}

var plcIO plcRegistors

type emissionCases struct {
	initial  []uint16 // initial emission values to start from (for 6 target)
	negative []uint16 // for negative samples, value should increase minimally or be same value for consecutive cycles
	positive []uint16 // for negative samples, increase over consecutive cycles
	testing  []uint16 // any other combination if you want per cycle like positive case with low load
}

var emissionCase = emissionCases{
	[]uint16{1100, 1100, 1100, 1100, 1100, 1100}, // initial
	[]uint16{0, 0, 0, 0, 0, 0},                   // negative
	[]uint16{10, 10, 10, 10, 10, 10},             // positive with high load
	[]uint16{10, 0, 5, 5, 0, 10},                 // positive with low load
}

const (
	roomTemp    float32 = 30 // assume room temp is 30.0 (300 for 30.0)
	wellsCount  uint16  = 96 // number of wells to simulate
	jitterValue int     = 50 // jitter to fluctuate emission value
)

type Simulator struct {
	ExitCh chan error
}

func NewSimulator(exit chan error) plc.Driver {
	return &Simulator{exit}
}

func (d *Simulator) HeartBeat() {
	logger.Info("Starting HeartBeat...")
	return
}

// ConfigureRun → Configure PCR Machine
func (d *Simulator) ConfigureRun(stg plc.Stage) error {
	plcIO.d.stage = stg
	fmt.Println("----------------------", plcIO.d.stage.IdealLidTemp)

	return nil
}

// Start → Starts the operations
func (d *Simulator) Start() error {
	time.Sleep(2 * time.Second)
	go simulate()
	return nil
}

// Monitor → periodically, if Status=CYCLE_COMPLETE, the Scan will be populated
func (d *Simulator) Monitor(cycle uint16) (scan plc.Scan, err error) {

	// Read current cycle
	scan.Cycle = 0

	// Read cycle temperature.. PLC returns 653 for 65.3 degrees
	scan.Temp = float32(plcIO.d.currentTemp) / 10

	// Read lid temperature
	scan.LidTemp = float32(plcIO.d.currentLidTemp) / 10

	// Read current cycle status
	tmp := 0

	if tmp == 0 { // 0x0000 means cycle is not complete
		// Values would not have changed.
		scan.CycleComplete = false
		return
	}
	scan.CycleComplete = true

	// If the invoker has already read this cycle data, don't send it again!
	if cycle == scan.Cycle {
		return
	}

	// Scan all the data from the Wells (96 x 6). Since max read is 123 registers, we shall read 96 at a time.
	//offset := 2000

	// for i := 0; i < 6; i++ {
	// 	var data []byte
	// 	data, err = C32.Driver.ReadHoldingRegisters(MODBUS["D"][offset+(i*96)], uint16(96))
	// 	if err != nil {
	// 		return
	// 	}

	// 	offset = 0 // offset of data. increment every 2 bytes!
	// 	for j := 0; j < 16; j++ {
	// 		// populate each wells with 6 emissions each
	// 		emission := plc.Emissions{}
	// 		for k := 0; k < 6; k++ {
	// 			emission[k] = binary.BigEndian.Uint16(data[offset : offset+2])
	// 			offset += 2
	// 		}

	// 		scan.Wells[(i*16)+j] = emission
	// 	}

	// }

	return
}

func simulate() {
	// simulatiing holding stage
	simulateHoldingStage()

	//simulate cycle stage
	//simulateCycleStage()
}

func simulateHoldingStage() {
	rt := roomTemp
	plcIO.d.currentTemp = uint16(rt * 10)

	for _, stp := range plcIO.d.stage.Holding {
		// ramping up temp
		for {
			// taking some time to increase the temperature
			time.Sleep(500 * time.Millisecond)

			// simulate currentLidTemp
			plcIO.d.currentLidTemp = jitter(uint16(plcIO.d.stage.IdealLidTemp*10), 1, 5)

			// simulate currentTemp
			plcIO.d.currentTemp = plcIO.d.currentTemp + uint16(stp.RampUpTemp*10)

			// if the target temp is below than the next multiple of ramp up temp
			if plcIO.d.currentTemp >= uint16(stp.TargetTemp*10) {
				plcIO.d.currentTemp = uint16(stp.TargetTemp * 10)
				break
			}

		}
	}
}

func simulateCycleStage(stg plc.Stage, ch chan plc.Scan) {
	scan := plc.Scan{}
	emissions := []plc.Emissions{}
	r := roomTemp

	for i := uint16(0); i < stg.CycleCount; i++ {
		scan.CycleComplete = false

		for s, stp := range stg.Cycle {
			for {
				scan.CycleComplete = false
				time.Sleep(time.Duration(stp.HoldTime) * time.Second) // spending time - HoldTime

				scan.Temp = r        // simulate cycle temp
				scan.LidTemp = r + 2 // lid temp is always a bit more than temp, ideally 2
				scan.Cycle = i + 1   // cycle is incrementing from 1

				if r >= stp.TargetTemp { // if the target temp is below than the next multiple of ramp up temp

					scan.Temp = stp.TargetTemp
					if s == len(stg.Cycle)-1 { // if last cycle
						scan.CycleComplete = true

						emissions = fillEmission(scan.Cycle, emissions) // populate emissions

						for x, vl := range emissions {
							scan.Wells[x] = vl
						}

					}
					ch <- scan

					break
				}
				ch <- scan

				r = r + stp.RampUpTemp
			}
			r = roomTemp
		}
		scan = plc.Scan{}
	}
}

func fillEmission(cycle uint16, ems []plc.Emissions) []plc.Emissions {

	emissions := []plc.Emissions{}
	emission := plc.Emissions{}

	for i := uint16(0); i < wellsCount; i++ {
		for x := range emission {
			emission[x] = jitter(emissionCase.initial[x], 1, 2)
		}

		if i < 31 { // first 32 wells are set for fail case
			for x := uint16(0); x < cycle-1; x++ {
				for i, v := range emission {
					emission[i] = v + jitter(emissionCase.negative[i], 1, 2)
				}
			}

			emissions = append(emissions, emission)
		}

		if i > 31 && i < 63 { // next 32 wells are set for pass case
			for x := uint16(0); x < cycle-1; x++ {
				for i, v := range emission {
					emission[i] = v + jitter(emissionCase.positive[i], 1, 2)
				}
			}

			emissions = append(emissions, emission)
		}

		if i > 63 && i < 95 { // next 32 wells are set for user-defined testing case
			for x := uint16(0); x < cycle-1; x++ {
				for i, v := range emission {
					emission[i] = v + jitter(emissionCase.testing[i], 1, 2)
				}
			}

			emissions = append(emissions, emission)
		}
	}

	return emissions
}

func jitter(n uint16, min, max int) uint16 {
	return n + uint16(rand.Intn((max-min))+min)
}
