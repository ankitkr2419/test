package service

import (
	"context"
	"fmt"
	"mylab/cpagent/db"
	"net/http"

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

		// Get the recipe
		recipe, err := deps.Store.ShowRecipe(req.Context(), recipeID)
		if err != nil {
			fmt.Fprintf(rw, err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		switch deck {
		case "A", "B":
			response, err = runRecipe(req.Context(), deps, deck, recipe)
		default:
			err = fmt.Errorf("Check you deck name")
		}

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

func runRecipe(ctx context.Context, deps Dependencies, deck string, recipe db.Recipe) (response string, err error) {

	// Get Processes associated with recipe
	processes, err := deps.Store.ListProcesses(ctx, recipe.ID)
	if err != nil {
		return "", err
	}

	currentCartridgeIDs := map[string]int64{
		"A": 0,
		"B": 0,
	}
	// No cartridge selected so cartridge_id by default is 0
	// Depending on cartridge_1 or cartridge_2 type we shall
	//  select cartridge_id from recipe field

	// currentTips will be maps of deck to TipsTubes
	currentTips := map[string]db.TipsTubes{
		"A": db.TipsTubes{},
		"B": db.TipsTubes{},
	}
	//  No tip selected
	//  This field will be set when a tip is picked up
	//  We will get its id from recipe and its details from tipsTubes map

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
				currentCartridgeIDs[deck] = recipe.Cartridge1Position
			} else {
				currentCartridgeIDs[deck] = recipe.Cartridge2Position
			}
			// TODO: Pass the complete Tip rather than just name for volume validations
			response, err = deps.PlcDeck[deck].AspireDispense(ad, currentCartridgeIDs[deck], currentTips[deck].Name)
			if err != nil {
				return "", err
			}
		case "Heating":
			heat, err := deps.Store.ShowHeating(ctx, p.ID)

			fmt.Printf("heat object %v", heat)
			ht, err := deps.PlcDeck[deck].Heating(uint16(heat.Temperature), heat.FollowTemp, heat.Duration)
			if err != nil {
				return "", err
			}
			fmt.Println(ht)

		case "Shaking":
			// Get the Shaking process
			// TODO: Below ID is reference ID, so please conform
			// sh, err := deps.Store.ShowShaking(req.Context(), p.ID)
			// if err != nil {
			// return "", err
			// }
			// sh.run()
		case "Piercing":
			// Get the Piercing process
			// TODO: Below ID is reference ID, so please conform
			pi, err := deps.Store.ShowPiercing(ctx, p.ID)
			if err != nil {
				return "", err
			}
			fmt.Println(pi)
			// pi.run()
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
				currentTips[deck], err = deps.Store.ShowTip(tipID)
				if err != nil {
					return "", err
				}
			case db.DiscardTip:
				currentTips[deck] = db.TipsTubes{}

			}
		case "TipDocking":
		case "Delay":

		}
		// TODO: Instead of switch case, try using reflect
		// Pass context and ID here
		// result := reflect.ValueOf(deps.PlcDeck[deck]).MethodByName(p.Type).Call([]reflect.Value{})
	}

	return
}

func getTipIDFromRecipePosition(recipe db.Recipe, position int64) (id int64, err error) {
	// Currently only 3 positions are allowed for tips
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
