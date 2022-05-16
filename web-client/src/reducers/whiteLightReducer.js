import { fromJS } from "immutable";
import {
  whiteLightActions,
  whiteLightDeckActions,
  whiteLightBothDeckActions,
} from "actions/whiteLightActions";
import { DECKNAME } from "appConstants";
import { getUpdatedDecks } from "utils/helpers";

export const lightInitialState = {
  decks: [
    {
      name: DECKNAME.DeckA,
      isLoading: false,
      isError: null,
      isLightOn: false,
    },
    {
      name: DECKNAME.DeckB,
      isLoading: false,
      isError: null,
      isLightOn: false,
    },
  ],
};

export const whiteLightReducer = (state = lightInitialState, action) => {
  switch (action.type) {
    case whiteLightDeckActions.initiateAction:
      const deckInitiateName =
        action.payload.params.deck === "A" ? DECKNAME.DeckA : DECKNAME.DeckB;

      const deckNumber = deckInitiateName == DECKNAME.DeckA ? 0 : 1;
      const changesForLightOn = {
        isLoading: false,
        isError: null,
        // isLightOn: !state.decks[deckNumber].isLightOn,
        isLightOn: action.payload.params.lightStatus,
      };

      const dockAfterRunInit = getUpdatedDecks(
        state,
        deckInitiateName,
        changesForLightOn
      );

      return {
        ...state,
        decks: dockAfterRunInit,
      };
    case whiteLightDeckActions.successAction:
      return {
        ...state,
      };
    case whiteLightDeckActions.failureAction:
      return {
        ...state,
      };
    case whiteLightDeckActions.resetAction:
      return {
        ...state,
      };
    default:
      return state;
  }
};
