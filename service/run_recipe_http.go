package service

import (
	"context"
	"fmt"
	"mylab/cpagent/db"
	"net/http"

	"github.com/google/uuid"

	"github.com/gorilla/mux"
)

func runRecipeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var response string
		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]

		recipeID, err := parseUUID(vars["id"])
		if err != nil {
			fmt.Fprintf(rw, err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		switch deck {
		case "A", "B":
			response, err = runRecipe(req.Context(), deps, deck, recipeID)
		default:
			err = fmt.Errorf("Check your deck name")
		}

		// TODO: Handle error types
		if err != nil {
			fmt.Fprintf(rw, err.Error())
			fmt.Println(err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
		} else {
			fmt.Fprintf(rw, response)
			rw.WriteHeader(http.StatusOK)
		}
	})
}

func runRecipe(ctx context.Context, deps Dependencies, deck string, recipeID uuid.UUID) (response string, err error) {

	if !deps.PlcDeck[deck].IsMachineHomed() {
		err = fmt.Errorf("Please home the machine first!")
		return
	}

	if deps.PlcDeck[deck].IsRunInProgress() {
		err = fmt.Errorf("previous run already in progress... wait or abort it")
		return
	}

	deps.PlcDeck[deck].SetRunInProgress()
	defer deps.PlcDeck[deck].ResetRunInProgress()

	// Get the recipe
	recipe, err := deps.Store.ShowRecipe(ctx, recipeID)
	if err != nil {
		return
	}

	// Get Processes associated with recipe
	processes, err := deps.Store.ListProcesses(ctx, recipe.ID)
	if err != nil {
		return "", err
	}

	var currentCartridgeID int64
	// No cartridge selected so cartridge_id by default is 0
	// Depending on cartridge_1 or cartridge_2 type we shall
	//  select cartridge_id from recipe field

	var currentTip db.TipsTubes
	//  No tip selected
	//  This field will be set when a tip is picked up
	//  We will get its id from recipe and its details from tipsTubes

	for _, p := range processes {
		switch p.Type {
		case "AspireDispense":
			// Get the AspireDispense process
			// TODO: Below ID is reference ID, so please change code accordingly
			ad, err := deps.Store.ShowAspireDispense(ctx, p.ID)
			if err != nil {
				return "", err
			}
			fmt.Println(ad)

			if ad.CartridgeType == db.Cartridge1 {
				currentCartridgeID = recipe.Cartridge1Position
			} else {
				currentCartridgeID = recipe.Cartridge2Position
			}
			// TODO: Pass the complete Tip rather than just name for volume validations
			response, err = deps.PlcDeck[deck].AspireDispense(ad, currentCartridgeID, currentTip.Name)
			if err != nil {
				return "", err
			}
		case "Heating":
			heat, err := deps.Store.ShowHeating(ctx, p.ID)
			fmt.Printf("heat object %v", heat)
			ht, err := deps.PlcDeck[deck].Heating(heat)
			if err != nil {
				return "", err
			}
			fmt.Println(ht)

		case "Shaking":
			// Get the Shaking process
			// sh, err := deps.Store.ShowShaking(req.Context(), p.ID)
			// if err != nil {
			// return "", err
			// }
			// fmt.Println(sh)
			// sh.run()
		case "Piercing":
			// Get the Piercing process
			// TODO: Below ID is reference ID, so please conform
			pi, err := deps.Store.ShowPiercing(ctx, p.ID)
			if err != nil {
				return "", err
			}
			fmt.Println(pi)

			if string(pi.Type) == db.Cartridge1 {
				currentCartridgeID = recipe.Cartridge1Position
			} else {
				currentCartridgeID = recipe.Cartridge2Position
			}

			response, err = deps.PlcDeck[deck].Piercing(pi, currentCartridgeID)
			if err != nil {
				return "", err
			}

		case "AttachDetach":
			ad, err := deps.Store.ShowAttachDetach(ctx, p.ID)
			fmt.Printf("attach detach record %v \n", ad)
			if err != nil {
				return "", err
			}
			response, err = deps.PlcDeck[deck].AttachDetach(ad)
			if err != nil {
				return "", err
			}

		case "TipOperation":
			to, err := deps.Store.ShowTipOperation(ctx, p.ID)
			if err != nil {
				return "", err
			}
			fmt.Println(to)

			response, err = deps.PlcDeck[deck].TipOperation(to)
			if err != nil {
				return "", err
			}

			switch to.Type {
			case db.PickupTip:
				// Store Current Tip here
				tipID, err := getTipIDFromRecipePosition(recipe, to.Position)
				if err != nil {
					return "", err
				}
				currentTip, err = deps.Store.ShowTip(tipID)
				if err != nil {
					return "", err
				}
			case db.DiscardTip:
				currentTip = db.TipsTubes{}

			}
		case "TipDocking":
			td, err := deps.Store.ShowTipDocking(ctx, p.ID)
			if err != nil {
				return "", err
			}
			fmt.Println(td)
			if td.Type == db.Cartridge1 {
				currentCartridgeID = recipe.Cartridge1Position
			} else {
				currentCartridgeID = recipe.Cartridge2Position
			}
			response, err = deps.PlcDeck[deck].TipDocking(td, currentCartridgeID)
			if err != nil {
				return "", err
			}
		case "Delay":
			delay, err := deps.Store.ShowDelay(ctx, p.ID)
			if err != nil {
				return "", err
			}
			fmt.Print(delay)
			response, err = deps.PlcDeck[deck].AddDelay(delay)
			if err != nil {
				return "", err
			}

		}
		// TODO: Instead of switch case, try using reflect
		// Pass context and ID here
		// result := reflect.ValueOf(deps.PlcDeck[deck]).MethodByName(p.Type).Call([]reflect.Value{})
	}

	// Home the machine
	deps.PlcDeck[deck].ResetRunInProgress()
	response, err = deps.PlcDeck[deck].Homing()
	if err != nil {
		return
	}

	return "SUCCESS", nil
}

func getTipIDFromRecipePosition(recipe db.Recipe, position int64) (id int64, err error) {
	// Currently only 3 positions are allowed for tips deck version 1.2
	// TODO: Change this for version 1.3
	switch position {
	case 1:
		return recipe.Position1, nil
	case 2:
		return recipe.Position2, nil
	case 3:
		return recipe.Position3, nil
	}
	err = fmt.Errorf("position is invalid to pickup the tip")
	return 0, err
}
