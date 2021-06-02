export const runRecipeAction = {
  runRecipeInitiated: "RUN_RECIPE_INITIATED",
  runRecipeSuccess: "RUN_RECIPE_SUCCESS",
  runRecipeFailed: "RUN_RECIPE_FAILED",
  runRecipeReset: "RUN_RECIPE_RESET",
  runRecipeInProgress: "PROGRESS_RUN_RECIPE",
  runRecipeInCompleted: "SUCCESS_RUN_RECIPE",
};

export const pauseRecipeAction = {
  pauseRecipeInitiated: "PAUSE_RECIPE_INITIATED",
  pauseRecipeSuccess: "PAUSE_RECIPE_SUCCESS",
  pauseRecipeFailed: "PAUSE_RECIPE_FAILED",
  pauseRecipeReset: "PAUSE_RECIPE_RESET",
};

export const resumeRecipeAction = {
  resumeRecipeInitiated: "RESUME_RECIPE_INITIATED",
  resumeRecipeSuccess: "RESUME_RECIPE_SUCCESS",
  resumeRecipeFailed: "RESUME_RECIPE_FAILED",
  resumeRecipeReset: "RESUME_RECIPE_RESET",
  resumeRecipeInProgress: "PROGRESS_RESUME_RECIPE",
  resumeRecipeInCompleted: "SUCCESS_RESUME_RECIPE",
};

export const abortRecipeAction = {
  abortRecipeInitiated: "ABORT_RECIPE_INITIATED",
  abortRecipeSuccess: "ABORT_RECIPE_SUCCESS",
  abortRecipeFailed: "ABORT_RECIPE_FAILED",
  abortRecipeReset: "ABORT_RECIPE_RESET",
};

export const recipeListingAction = {
  recipeListingInitiated: "RECIPE_LISTING_INITIATED",
  recipeListingSuccess: "RECIPE_LISTING_SUCSESSS",
  recipeListingFailed: "RECIPE_LISTING_FAILED",
};

export const saveRecipeDataAction = {
  saveRecipeDataForDeck: "SAVE_RECIPE_DATA_FOR_DECK",
  resetRecipeDataForDeck: "RESET_RECIPE_DATA_FOR_DECK",
  updateRecipeReducerDataForDeck: "UPDATE_RECIPE_REDUCER_DATA_FOR_DECK"
}

export const stepRunRecipeAction = {
  stepRunRecipeInitiated: "STEP_RUN_RECIPE_INITIATED",
  stepRunRecipeSuccess: "RUN_RECIPE_SUCCESS",
  stepRunRecipeFailed: "RUN_RECIPE_FAILED",
  nextStepRunRecipeInitiated: "NEXT_STEP_RUN_RECIPE_INITIATED",
}

export const publishRecipeAction = {
  publishRecipeInitiated: "PUBLISH_RECIPE_INITIATED",
  publishRecipeSuccess: "PUBLISH_RECIPE_SUCCESS",
  publishRecipeFailed: "PUBLISH_RECIPE_FAILED"
}

export const deleteRecipeAction = {
  deleteRecipeInitiated: "DELETE_RECIPE_INITIATED",
  deleteRecipeSuccess: "DELETE_RECIPE_SUCCESS",
  deleteRecipeFailure: "DELETE_RECIPE_FAILURE",
}