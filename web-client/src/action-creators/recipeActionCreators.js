import {
  runRecipeAction,
  pauseRecipeAction,
  resumeRecipeAction,
  abortRecipeAction,
  recipeListingAction,
  saveRecipeDataAction,
  stepRunRecipeAction,
  publishRecipeAction,
  deleteRecipeAction,
} from "actions/recipeActions";

export const runRecipeInitiated = (params) => ({
  type: runRecipeAction.runRecipeInitiated,
  payload: {
    params,
  },
});

export const runRecipeSuccess = (runRecipeResponse) => ({
  type: runRecipeAction.runRecipeSuccess,
  payload: {
    runRecipeResponse,
  },
});

export const runRecipeFailed = (serverErrors) => ({
  type: runRecipeAction.runRecipeFailed,
  payload: {
    serverErrors,
  },
});

export const runRecipeReset = (deckName) => ({
  type: runRecipeAction.runRecipeReset,
  payload: { deckName },
});

export const runRecipeInProgress = (runRecipeInProgress) => ({
  type: runRecipeAction.runRecipeInProgress,
  payload: {
    runRecipeInProgress,
  },
});

export const runRecipeInCompleted = (runRecipeInCompleted) => ({
  type: runRecipeAction.runRecipeInCompleted,
  payload: {
    runRecipeInCompleted,
  },
});

export const pauseRecipeInitiated = (params) => ({
  type: pauseRecipeAction.pauseRecipeInitiated,
  payload: {
    params,
  },
});

export const pauseRecipeSuccess = (pauseRecipeResponse) => ({
  type: pauseRecipeAction.pauseRecipeSuccess,
  payload: {
    pauseRecipeResponse,
  },
});

export const pauseRecipeFailed = (serverErrors) => ({
  type: pauseRecipeAction.pauseRecipeFailed,
  payload: {
    serverErrors,
  },
});

export const pauseRecipeReset = () => ({
  type: pauseRecipeAction.pauseRecipeReset,
});

export const resumeRecipeInitiated = (params) => ({
  type: resumeRecipeAction.resumeRecipeInitiated,
  payload: {
    params,
  },
});

export const resumeRecipeSuccess = (resumeRecipeResponse) => ({
  type: resumeRecipeAction.resumeRecipeSuccess,
  payload: {
    resumeRecipeResponse,
  },
});

export const resumeRecipeFailed = (serverErrors) => ({
  type: resumeRecipeAction.resumeRecipeFailed,
  payload: {
    serverErrors,
  },
});

export const resumeRecipeReset = () => ({
  type: resumeRecipeAction.resumeRecipeReset,
});

export const resumeRecipeInProgress = (resumeRecipeInProgress) => ({
  type: resumeRecipeAction.resumeRecipeInProgress,
  payload: resumeRecipeInProgress,
});

export const resumeRecipeInCompleted = (resumeRecipeInCompleted) => ({
  type: resumeRecipeAction.resumeRecipeInCompleted,
  payload: resumeRecipeInCompleted,
});

export const abortRecipeInitiated = (params) => ({
  type: abortRecipeAction.abortRecipeInitiated,
  payload: {
    params,
  },
});

export const abortRecipeSuccess = (abortRecipeResponse) => ({
  type: abortRecipeAction.abortRecipeSuccess,
  payload: {
    abortRecipeResponse,
  },
});

export const abortRecipeFailed = (serverErrors) => ({
  type: abortRecipeAction.abortRecipeFailed,
  payload: {
    serverErrors,
  },
});

export const abortRecipeReset = () => ({
  type: abortRecipeAction.abortRecipeReset,
});

export const recipeListingInitiated = (token, deckName) => ({
  type: recipeListingAction.recipeListingInitiated,
  payload: {
    token,
    deckName,
  },
});

export const recipeListingSuccess = (recipeData) => ({
  type: recipeListingAction.recipeListingSuccess,
  payload: {
    recipeData,
  },
});

export const recipeListingFailed = (serverErrors) => ({
  type: recipeListingAction.recipeListingFailed,
  payload: {
    serverErrors,
  },
});

export const saveRecipeDataForDeck = (recipeData, deckName) => ({
  //deckName should be passed
  type: saveRecipeDataAction.saveRecipeDataForDeck,
  payload: {
    recipeData,
    deckName,
  },
});

export const resetRecipeDataForDeck = (deckName) => ({
  //deckName should be passed
  type: saveRecipeDataAction.resetRecipeDataForDeck,
  payload: {
    deckName,
  },
});

export const updateRecipeReducerDataForDeck = (deckName, body) => ({
  //deckName should be passed
  type: saveRecipeDataAction.updateRecipeReducerDataForDeck,
  payload: {
    deckName,
    body,
  },
});

export const stepRunRecipeInitiated = (params) => ({
  type: stepRunRecipeAction.stepRunRecipeInitiated,
  payload: {
    params,
  },
});

export const nextStepRunRecipeInitiated = (params) => ({
  type: stepRunRecipeAction.nextStepRunRecipeInitiated,
  payload: {
    params,
  },
});

export const publishRecipeInitiated = (params) => ({
  type: publishRecipeAction.publishRecipeInitiated,
  payload: {
    params,
  },
});

export const deleteRecipeInitiated = (params) => ({
  type: deleteRecipeAction.deleteRecipeInitiated,
  payload: {
    params,
  },
});
