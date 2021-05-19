import { DECKNAME } from "appConstants";
import {
  getTipsAndTubesAction,
  getCartridgeAction,
  saveNewRecipeAction,
  getTubesAction,
  getTipsAction,
} from "actions/saveNewRecipeActions";
// import { fromJS } from "immutable";

//  Old init state

// const initialState = {
//   tempDeckName: "",
//   isLoading: null,
//   error: null,
//   isSuccess: null,
//   decks: [
//     {
//       name: DECKNAME.DeckA,
//       recipeDetails: {
//         name: "",
//       },
//       recipeOptions: null,
//       tubesOptions: null,
//       tipsOptions: null,
//       cartridgeOptions: null,
//       isSaved: false,
//       errorInSaving: false,
//       token: "",
//     },
//     {
//       name: DECKNAME.DeckB,
//       recipeDetails: {
//         name: "",
//       },
//       recipeOptions: null,
//       tubesOptions: null,
//       tipsOptions: null,
//       cartrideOptions: null,
//       isSaved: false,
//       errorInSaving: false,
//       token: "",
//     },
//   ],
// };

const initialState = {
  tempDeckName: "",
  isLoading: null,
  error: null,
  isSuccess: null,
  name: DECKNAME.DeckA,
  recipeDetails: {
    name: "",
  },
  recipeOptions: null,
  tubesOptions: null,
  tipsOptions: null,
  cartridgeOptions: null,
  isSaved: false,
  errorInSaving: false,
  token: "",
};

export const updateRecipeDetailsReducer = (state = initialState, actions) => {
  switch (actions.type) {
    //saving recipe name
    case saveNewRecipeAction.saveRecipeName:
      return {
        ...state.recipeDetails,
        ...actions.payload.recipeDetails,
      };

    //saving and updating new recipe
    case saveNewRecipeAction.updateRecipeInitiated:
      const deckName = actions.payload.deckName;
      const token = actions.payload.token;

      return {
        ...state,
        isLoading: true,
        tempDeckName: deckName,
        token: token,
      };

    case saveNewRecipeAction.updateRecipeSuccess:
      return { ...state, isLoading: false, error: false, isSuccess: true };

    case saveNewRecipeAction.updateRecipeFailure:
      return {
        ...state,
        isLoading: false,
        error: true,
        isSuccess: false,
      };

    case saveNewRecipeAction.updateRecipeReset:
      return {
        ...state,
        isLoading: null,
        error: null,
        isSuccess: null,
      };

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

    //tubes options
    case getTubesAction.getTubesInitiated:
      return {
        ...state,
        isLoading: true,
        tempDeckName: actions.payload.deckName,
      };

    case getTubesAction.getTubesSuccess:
      let deckAfterTubesLoadingSuccess = state.decks.map((deckObj, index) => {
        return deckObj.name === state.tempDeckName
          ? {
              ...deckObj,
              tubesOptions: actions.payload.response,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: deckAfterTubesLoadingSuccess,
        isLoading: false,
        error: false,
      };

    case getTubesAction.getTubesFailure:
      return {
        ...state,
        isLoading: false,
        error: true,
      };

    case getTubesAction.getTubesReset:
      return {
        ...state,
        isLoading: null,
        error: null,
      };

    //tips options
    case getTipsAction.getTipsInitiated:
      return {
        ...state,
        isLoading: true,
        tempDeckName: actions.payload.deckName,
      };

    case getTipsAction.getTipsSuccess:
      let deckAfterTipsLoadingSuccess = state.decks.map((deckObj, index) => {
        return deckObj.name === state.tempDeckName
          ? {
              ...deckObj,
              tipsOptions: actions.payload.response,
            }
          : deckObj;
      });
      return {
        ...state,
        decks: deckAfterTipsLoadingSuccess,
        isLoading: false,
        error: false,
      };

    case getTipsAction.getTipsFailure:
      return {
        ...state,
        isLoading: false,
        error: true,
      };

    case getTipsAction.getTipsReset:
      return {
        ...state,
        isLoading: null,
        error: null,
      };

    default:
      return state;
  }
};
