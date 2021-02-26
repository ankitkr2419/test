package service

import (
	"context"
	"fmt"
	"mylab/cpagent/db"
	"net/http"
	"time"

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

		// Get Processes associated with recipe
		processes, err := deps.Store.ListProcesses(req.Context(), recipeID)
		if err != nil {
			fmt.Fprintf(rw, err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		switch deck {
		case "":
			response, err = runRecipeOnBothDeck(req.Context(), deps, deck, processes, recipe)
		case "A", "B":
			response, err = runRecipe(req.Context(), deps, deck, processes, recipe)
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

func runRecipe(ctx context.Context, deps Dependencies, deck string, processes []db.Process, recipe db.Recipe) (response string, err error) {

	// No cartridge selected so cartridge_id by default is 0
	// Depending on cartridge_1 or cartridge_2 type we shall
	//  select cartridge_id from recipe field
	// var cartridgeID int64 = 0
	//  No tip selected
	//  This field will be set when a tip is picked up
	//  We will get its id from recipe and its details from tipsTubes map
	// var tipType string = ""

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
			// Call Deck Process here
		case "Heating":
			// Get the Heating process
			// TODO: Below ID is reference ID, so please conform
			// ht, err := deps.Store.ShowHeating(req.Context(), p.ID)
			// if err != nil {
			// return "", err
			// }
		case "Shaking":
			// Get the Shaking process
			// TODO: Below ID is reference ID, so please conform
			// sh, err := deps.Store.ShowShaking(req.Context(), p.ID)
			// if err != nil {
			// return "", err
			// }
		case "Piercing":
			// Get the Piercing process
			// TODO: Below ID is reference ID, so please conform
			pi, err := deps.Store.ShowPiercing(ctx, p.ID)
			if err != nil {
				return "", err
			}
			fmt.Println(pi)
		case "Magnet":
		case "TipOperation":
		case "TipDocking":
		case "Delay":

		}
		// TODO: Instead of switch case, try using reflect
		// Pass context and ID here
		// result := reflect.ValueOf(deps.PlcDeck[deck]).MethodByName(p.Type).Call([]reflect.Value{})
	}

	return
}

func runRecipeOnBothDeck(ctx context.Context, deps Dependencies, deck string, processes []db.Process, recipe db.Recipe) (response string, err error) {

	var deckAResponse, deckBResponse string
	var deckAErr, deckBErr error

	//  If we need both recipes to go at the same pace
	// then delegate the deck chossing till before fetching invidual process
	go func() {
		deckAResponse, deckAErr = runRecipe(ctx, deps, "A", processes, recipe)

	}()
	go func() {
		deckBResponse, deckBErr = runRecipe(ctx, deps, "B", processes, recipe)
	}()

	for {
		switch {
		case deckAErr != nil:
			fmt.Printf("Error deck A %v", deckAErr)
			return "", deckAErr
		case deckBErr != nil:
			fmt.Printf("Error deck B %v", deckBErr)
			return "", deckBErr
		case deckAResponse != "" && deckBResponse != "":
			operationSuccessMsg := fmt.Sprintf("Success for both Decks!")
			fmt.Println(operationSuccessMsg)
			return operationSuccessMsg, nil
		default:
			// Only check every 400 milli second
			time.Sleep(400 * time.Millisecond)
		}
	}
}
