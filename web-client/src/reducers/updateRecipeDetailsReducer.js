import { DECKNAME } from "appConstants";
import {
  getRecipeDetailsAction,
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
      isSaved: false,
      errorInSaving: false,
    },
    {
      name: DECKNAME.DeckB,
      recipeDetails: {
        name: "",
      },
      recipeOptions: null,
      isSaved: false,
      errorInSaving: false,
    },
  ],
};

export const updateRecipeDetailsReducer = (state = initialState, actions) => {
  switch (actions.type) {
    case saveNewRecipeAction.saveRecipeName:
      const deckNameToSaveRecipeTo = actions.payload.deckName;
      let deckAfterSave = state.decks.map((deckObj, index) => {
        return deckObj.name === deckNameToSaveRecipeTo
          ? {
              ...deckObj,
              recipeDetails: {
                ...deckObj.recipeDetails,
                ...actions.payload.recipeDetails,
              },
            }
          : deckObj;
      });
      return {
        ...state,
        decks: deckAfterSave,
      };

    //Save options
    case getRecipeDetailsAction.getRecipeDetailsInitiated:
      return {
        ...state,
        isLoading: true,
        tempDeckName: actions.payload.deckName,
      };

    case getRecipeDetailsAction.getRecipeDetailsSuccess:
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

    case getRecipeDetailsAction.getRecipeDetailsFailure:
      return {
        ...state,
        isLoading: false,
        error: true,
      };

    case getRecipeDetailsAction.getRecipeDetailsReset:
      return {
        ...state,
        isLoading: null,
        error: null,
      };

    default:
      return state;
  }
};

// const saveRecipeInitialState = fromJS({
//   name: "",
//   //isSaved
//   //deckName
//   //others...
// });

// export const saveNewRecipeReducer = (
//   state = saveRecipeInitialState,
//   action
// ) => {
//   switch (action.type) {
//     case saveNewRecipeActions.updateRecipe:
//       return state.merge({
//         ...action.payload,
//       });
//     case saveNewRecipeActions.resetRecipe:
//       return saveRecipeInitialState;
//     default:
//       return state;
//   }
// };
