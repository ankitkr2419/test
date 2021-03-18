import {
  runRecipeAction,
  pauseRecipeAction,
  resumeRecipeAction,
  abortRecipeAction,
  recipeListingAction
} from "actions/recipeActions";

export const runRecipeInitiated = (params) => ({
  type: runRecipeAction.runRecipeInitiated,
  payload: {
    params
  }
});

export const runRecipeSuccess = (runRecipeResponse) => ({
  type: runRecipeAction.runRecipeSuccess,
  payload: {
    runRecipeResponse
  }
});

export const runRecipeFailed = (serverErrors) => ({
  type: runRecipeAction.runRecipeFailed,
  payload: {
    serverErrors
  }
});

export const pauseRecipeInitiated = (params) => ({
  type: pauseRecipeAction.pauseRecipeInitiated,
  payload: {
    params
  }
});

export const pauseRecipeSuccess = (pauseRecipeResponse) => ({
  type: pauseRecipeAction.pauseRecipeSuccess,
  payload: {
    pauseRecipeResponse
  }
});

export const pauseRecipeFailed = (serverErrors) => ({
  type: pauseRecipeAction.pauseRecipeFailed,
  payload: {
    serverErrors
  }
});

export const resumeRecipeInitiated = (params) => ({
  type: resumeRecipeAction.resumeRecipeInitiated,
  payload: {
    params
  }
});

export const resumeRecipeSuccess = (resumeRecipeResponse) => ({
  type: resumeRecipeAction.resumeRecipeSuccess,
  payload: {
    resumeRecipeResponse
  }
});

export const resumeRecipeFailed = (serverErrors) => ({
  type: resumeRecipeAction.resumeRecipeFailed,
  payload: {
    serverErrors
  }
});

export const abortRecipeInitiated = (params) => ({
  type: abortRecipeAction.abortRecipeInitiated,
  payload: {
    params
  }
});

export const abortRecipeSuccess = (abortRecipeResponse) => ({
  type: abortRecipeAction.abortRecipeSuccess,
  payload: {
    abortRecipeResponse
  }
});

export const abortRecipeFailed = (serverErrors) => ({
  type: abortRecipeAction.abortRecipeFailed,
  payload: {
    serverErrors
  }
});

export const recipeListingInitiated = () => ({
  type: recipeListingAction.recipeListingInitiated,
  payload: { }
});

export const recipeListingSuccess = (recipeData) => ({
  type: recipeListingAction.recipeListingSuccess,
  payload: {
    recipeData
  }
});

export const recipeListingFailed = (serverErrors) => ({
  type: recipeListingAction.recipeListingFailed,
  payload: {
    serverErrors
  }
});
