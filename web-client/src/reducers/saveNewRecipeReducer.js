import { DECKNAME } from "appConstants";
import { saveNewRecipeAction } from "actions/saveNewRecipeActions";
// import { fromJS } from "immutable";

const initialState = {
  decks: [
    {
      name: DECKNAME.DeckA,
      recipeName: "",
      isLoading: false,
      isSaved: false,
      errorInSaving: false,
    },
    {
      name: DECKNAME.DeckB,
      recipeName: "",
      isLoading: false,
      isSaved: false,
      errorInSaving: false,
    },
  ],
};

export const saveNewRecipeReducer = (state = initialState, actions) => {
  switch (actions.type) {
    case saveNewRecipeAction.saveRecipeName:
      const deckNameToSaveRecipeTo = actions.payload.deckName;
      let deckAfterSave = state.decks.map((deckObj, index) => {
        return deckObj.name === deckNameToSaveRecipeTo
          ? {
              ...state.decks.find(
                (initialDeckObj) =>
                  initialDeckObj.name === deckNameToSaveRecipeTo
              ),
              recipeName: actions.payload.recipeName,
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
