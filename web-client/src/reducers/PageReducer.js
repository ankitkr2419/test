import { pageBtnAction } from "actions/PageActions";

export const initialState = {
  recipePageDeckA: 1,
  recipePageDeckB: 1,
  processPageDeckA: 1,
  processPageDeckB: 1,
};

export const pageReducer = (state = initialState, action = {}) => {
  switch (action.type) {
    case pageBtnAction.nextRecipePageBtn:
      const deckNameRecipeNextPage = action.payload;
      if (deckNameRecipeNextPage === "Deck A") {
        return {
          ...state,
          recipePageDeckA: state.recipePageDeckA + 1,
        };
      } else {
        return {
          ...state,
          recipePageDeckB: state.recipePageDeckB + 1,
        };
      }
    case pageBtnAction.previousRecipePageBtn:
      const deckNameRecipePrevPage = action.payload;
      if (deckNameRecipePrevPage === "Deck A") {
        return {
          ...state,
          recipePageDeckA: state.recipePageDeckA - 1,
        };
      } else {
        return {
          ...state,
          recipePageDeckB: state.recipePageDeckB - 1,
        };
      }
    case pageBtnAction.nextProcessPageBtn:
      const deckNameProcessNextPage = action.payload;
      if (deckNameProcessNextPage === "Deck A") {
        return {
          ...state,
          processPageDeckA: state.processPageDeckA + 1,
        };
      } else {
        return {
          ...state,
          processPageDeckB: state.processPageDeckB + 1,
        };
      }
    case pageBtnAction.previousProcessPageBtn:
      const deckNameProcessPrevPage = action.payload;
      if (deckNameProcessPrevPage === "Deck A") {
        return {
          ...state,
          processPageDeckA: state.processPageDeckA - 1,
        };
      } else {
        return {
          ...state,
          processPageDeckB: state.processPageDeckB - 1,
        };
      }
    case pageBtnAction.ResetRecipePageState:
      const deckNameRecipeReset = action.payload;
      if (deckNameRecipeReset === "Deck A") {
        return {
          ...state,
          recipePageDeckA: 1,
        };
      } else {
        return {
          ...state,
          recipePageDeckB: 1,
        };
      }
    case pageBtnAction.ResetProcessPageState:
      const deckNameProcessReset = action.payload;
      if (deckNameProcessReset === "Deck A") {
        return {
          ...state,
          processPageDeckA: 1,
        };
      } else {
        return {
          ...state,
          processPageDeckB: 1,
        };
      }
    default:
      return state;
  }
};
