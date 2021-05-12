import { DECKNAME } from "appConstants";
import {
  getTipsAndTubesAction,
  getCartridgeAction,
  saveNewRecipeAction,
} from "actions/saveNewRecipeActions";
// import { fromJS } from "immutable";

const initialState = {
  tempDeckName: "",
  isLoading: null,
  error: null,
  decks: [
    {
      name: DECKNAME.DeckA,
      recipeDetails: {
        name: "",
      },
      recipeOptions: null,
      cartridgeOptions: null,
      isSaved: false,
      errorInSaving: false,
      token: "",
    },
    {
      name: DECKNAME.DeckB,
      recipeDetails: {
        name: "",
      },
      recipeOptions: null,
      cartrideOptions: null,
      isSaved: false,
      errorInSaving: false,
      token: "",
    },
  ],
};

export const updateRecipeDetailsReducer = (state = initialState, actions) => {
  switch (actions.type) {
    //saving and updating new recipe
    case saveNewRecipeAction.updateRecipeInitiated:
      const requestBody = actions.payload.requestBody;
      const deckName = actions.payload.deckName;

      const deckNameToSaveRecipeTo = deckName;
      return {
        ...state,
        isLoading: true,
        tempDeckName: deckName,
      };

    case saveNewRecipeAction.updateRecipeSuccess:
      console.log(actions);

      return state;

    // const deckNameToSaveRecipeTo = deckName;
    // return {
    //   ...state,
    //   isLoading: true,
    //   tempDeckName: deckName,
    // };

    // let deckAfterSave = state.decks.map((deckObj, index) => {
    //   return deckObj.name === deckNameToSaveRecipeTo
    //     ? {
    //         ...deckObj,
    //         recipeDetails: {
    //           ...deckObj.recipeDetails,
    //           ...actions.payload.recipeDetails,
    //         },
    //       }
    //     : deckObj;
    // });
    // return {
    //   ...state,
    //   decks: deckAfterSave,
    // };

    //tips and tubes options
    case getTipsAndTubesAction.getTipsAndTubesInitiated:
      return {
        ...state,
        isLoading: true,
        tempDeckName: actions.payload.deckName,
      };

    case getTipsAndTubesAction.getTipsAndTubesSuccess:
      let deckAfterLoadingSuccess = state.decks.map((deckObj, index) => {
        return deckObj.name === state.tempDeckName
          ? {
              ...deckObj,
              recipeOptions: actions.payload.response,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: deckAfterLoadingSuccess,
        isLoading: false,
        error: false,
      };

    case getTipsAndTubesAction.getTipsAndTubesFailure:
      return {
        ...state,
        isLoading: false,
        error: true,
      };

    case getTipsAndTubesAction.getTipsAndTubesReset:
      return {
        ...state,
        isLoading: null,
        error: null,
      };

    //cartridge options
    case getCartridgeAction.getCartridgeInitiated:
      return {
        ...state,
        isLoading: true,
        tempDeckName: actions.payload.deckName,
      };

    case getCartridgeAction.getCartridgeSuccess:
      let deckAfterCartridgeLoadingSuccess = state.decks.map(
        (deckObj, index) => {
          return deckObj.name === state.tempDeckName
            ? {
                ...deckObj,
                cartridgeOptions: actions.payload.response,
              }
            : deckObj;
        }
      );
      return {
        ...state,
        decks: deckAfterCartridgeLoadingSuccess,
        isLoading: false,
        error: false,
      };

    case getCartridgeAction.getCartridgeFailure:
      return {
        ...state,
        isLoading: false,
        error: true,
      };

    case getCartridgeAction.getCartridgeReset:
      return {
        ...state,
        isLoading: null,
        error: null,
      };

    default:
      return state;
  }
};
