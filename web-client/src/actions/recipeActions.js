export const runRecipeAction = {
  runRecipeInitiated: "RUN_RECIPE_INITIATED",
  runRecipeSuccess: "RUN_RECIPE_SUCCESS",
  runRecipeFailed: "RUN_RECIPE_FAILED",
  runRecipeReset: "RUN_RECIPE_RESET",
  runRecipeInProgress: "PROGRESS_RUN_RECIPE",
  runRecipeInCompleted: "SUCCESS_RUN_RECIPE"
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
  resumeRecipeInCompleted: "SUCCESS_RESUME_RECIPE"
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
  recipeListingReset: "RECIPE_LISTING_RESET",
};
