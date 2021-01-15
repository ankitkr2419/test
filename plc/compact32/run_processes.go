package compact32

import (
	"fmt"
)

func (d *Compact32Deck) RunProcesses() (response string, err error) {

	var category, cartridgeType string
	var labwareID, source_well, destination_well, aspire_cycles, dispense_cycles int64
	var asp_height, asp_mix_vol, asp_vol, dis_height, dis_mix_vol, dis_vol, dis_blow float64

	// Only extractions
	// 1. well_to_well
	category = "well_to_well"
	cartridgeType = "extraction"

	labwareID = 1
	source_well = 1
	destination_well = 2
	aspire_cycles = 5
	dispense_cycles = 3

	asp_height = 10
	asp_mix_vol = 100
	asp_vol = 120
	dis_height = 10
	dis_mix_vol = 100
	dis_vol = 120
	dis_blow = 10

	response, err = d.AspireDispense(category, cartridgeType,
		labwareID, source_well, destination_well, aspire_cycles, dispense_cycles,
		asp_height, asp_mix_vol, asp_vol, dis_height, dis_mix_vol, dis_vol, dis_blow)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("Problem doing aspire/dispense")
	}

	// 2. well_to_shaker
	category = "well_to_shaker"
	response, err = d.AspireDispense(category, cartridgeType,
		labwareID, source_well, destination_well, aspire_cycles, dispense_cycles,
		asp_height, asp_mix_vol, asp_vol, dis_height, dis_mix_vol, dis_vol, dis_blow)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("Problem doing aspire/dispense")
	}

	// 3. shaker_to_well
	category = "shaker_to_well"
	response, err = d.AspireDispense(category, cartridgeType,
		labwareID, source_well, destination_well, aspire_cycles, dispense_cycles,
		asp_height, asp_mix_vol, asp_vol, dis_height, dis_mix_vol, dis_vol, dis_blow)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("Problem doing aspire/dispense")
	}

	return "RUN Processes success", nil
}
