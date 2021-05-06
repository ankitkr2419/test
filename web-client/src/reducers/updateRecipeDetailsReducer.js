import { DECKNAME } from "appConstants";
import { saveNewRecipeAction } from "actions/saveNewRecipeActions";
// import { fromJS } from "immutable";

const initialState = {
  decks: [
    {
      name: DECKNAME.DeckA,
      recipeDetails: {
        name: "",
      },
      isLoading: false,
      isSaved: false,
      errorInSaving: false,
    },
    {
      name: DECKNAME.DeckB,
      recipeDetails: {
        name: "",
      },
      isLoading: false,
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
