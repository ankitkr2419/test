export const saveNewRecipeAction = {
  saveRecipeInitiated: "SAVE_RECIPE_INITIATE",
  saveRecipeName: "SAVE_RECIPE_NAME",
  saveRecipeSuccess: "SAVE_RECIPE_SUCCESS",
  saveRecipeFailed: "SAVE_RECIPE_FAILED",
  saveRecipeReset: "SAVE_RECIPE_RESET",
};

export const getRecipeDetailsAction = {
  getRecipeDetailsInitiated: "GET_RECIPE_DETAILS_INITIATED",
  getRecipeDetailsSuccess: "GET_RECIPE_DETAILS_SUCCESS",
  getRecipeDetailsFailure: "GET_RECIPE_DETAILS_FAILED",
  getRecipeDetailsReset: "GET_RECIPE_DETAILS_RESET",
}

export default {
  updateRecipe: "UPDATE_RECIPE",
  saveRecipe: "SAVE_RECIPE",
  resetRecipe: "RESET_RECIPE",
};
