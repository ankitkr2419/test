import { runRecipeAction, pauseRecipeAction, resumeRecipeAction, abortRecipeAction, recipeListingAction } from "actions/recipeActions";

export const initialState = {
  isLoading: false,
  serverErrors: {},
  runRecipeResponse: {},
  abortRecipeResponse: {},
  pauseRecipeResponse: {},
  resumeRecipeResponse: {},
  recipeData: {}
};

export const recipeActionReducer = (state = initialState, action = {}) => {
  switch (action.type) {
    case runRecipeAction.runRecipeInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case runRecipeAction.runRecipeSuccess:
      return { ...state, ...action.payload, isLoading: false };
    case runRecipeAction.runRecipeFailed:
      return { ...state, ...action.payload, isLoading: false };
    case pauseRecipeAction.pauseRecipeInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case pauseRecipeAction.pauseRecipeSuccess:
      return { ...state, ...action.payload, isLoading: false };
    case pauseRecipeAction.pauseRecipeFailed:
      return { ...state, ...action.payload, isLoading: false };
    case abortRecipeAction.abortRecipeInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case abortRecipeAction.abortRecipeSuccess:
      return { ...state, ...action.payload, isLoading: false };
    case abortRecipeAction.abortRecipeFailed:
      return { ...state, ...action.payload, isLoading: false };
    case resumeRecipeAction.resumeRecipeInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case resumeRecipeAction.resumeRecipeSuccess:
      return { ...state, ...action.payload, isLoading: false };
    case resumeRecipeAction.resumeRecipeFailed:
      return { ...state, ...action.payload, isLoading: false };
    case recipeListingAction.recipeListingInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case recipeListingAction.recipeListingSuccess:
      return { ...state, ...action.payload, isLoading: false };
    case recipeListingAction.recipeListingFailed:
      return { ...state, ...action.payload, isLoading: false };
    default:
      return state;
  }
}