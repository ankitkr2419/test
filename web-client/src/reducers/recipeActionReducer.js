import {
  runRecipeAction,
  pauseRecipeAction,
  resumeRecipeAction,
  abortRecipeAction,
  recipeListingAction,
} from "actions/recipeActions";
import { DECKCARD_BTN } from "appConstants";

export const initialState = {
  isLoading: false,
  serverErrors: {},
  runRecipeError: false,
  abortRecipeError: false,
  pauseRecipeError: false,
  resumeRecipeError: false,
  recipeListingError: false,
  runRecipeResponse: {},
  abortRecipeResponse: {},
  pauseRecipeResponse: {},
  resumeRecipeResponse: {},
  recipeData: [],
  leftActionBtn: DECKCARD_BTN.text.run,
  rightActionBtn: DECKCARD_BTN.text.cancel,
};

export const recipeActionReducer = (state = initialState, action = {}) => {
  switch (action.type) {
    case runRecipeAction.runRecipeInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case runRecipeAction.runRecipeSuccess:
      return {
        ...state,
        runRecipeResponse: action.payload.response,
        isLoading: false,
        leftActionBtn: DECKCARD_BTN.text.pause,
        rightActionBtn: DECKCARD_BTN.text.abort,
      };
    case runRecipeAction.runRecipeFailed:
      return {
        ...state,
        serverErrors: action.payload.serverErrors,
        isLoading: false,
        runRecipeError: true,
      };

    case pauseRecipeAction.pauseRecipeInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case pauseRecipeAction.pauseRecipeSuccess:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        leftActionBtn: DECKCARD_BTN.text.resume,
      };
    case pauseRecipeAction.pauseRecipeFailed:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        pauseRecipeAction: true,
      };

    case abortRecipeAction.abortRecipeInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case abortRecipeAction.abortRecipeSuccess:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        leftActionBtn: DECKCARD_BTN.text.run,
        rightActionBtn: DECKCARD_BTN.text.cancel,
      };
    case abortRecipeAction.abortRecipeFailed:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        abortRecipeError: true,
      };

    case resumeRecipeAction.resumeRecipeInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case resumeRecipeAction.resumeRecipeSuccess:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        leftActionBtn: DECKCARD_BTN.text.pause,
      };
    case resumeRecipeAction.resumeRecipeFailed:
      return {
        ...state,
        ...action.payload,
        isLoading: false,
        resumeRecipeError: true,
      };

    case recipeListingAction.recipeListingInitiated:
      return { ...state, ...action.payload, isLoading: true };
    case recipeListingAction.recipeListingSuccess:
      return {
        ...state,
        recipeData: action.payload.response,
        isLoading: false,
      };
    case recipeListingAction.recipeListingFailed:
      return {
        ...state,
        serverErrors: action.payload.serverErrors,
        error: true,
        isLoading: false,
        recipeListingError: true,
      };

    default:
      return state;
  }
};
