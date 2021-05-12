export const saveNewRecipeAction = {
  saveRecipeName: "SAVE_RECIPE_NAME",
  updateRecipeInitiated: "UPDATE_RECIPE_INITIATE",
  updateRecipeSuccess: "UPDATE_RECIPE_SUCCESS",
  updateRecipeFailure: "UPDATE_RECIPE_FAILED",
  updateRecipeReset: "UPDATE_RECIPE_RESET",
};

export const getTipsAndTubesAction = {
  getTipsAndTubesInitiated: "GET_TIPS_AND_TUBES_INITIATED",
  getTipsAndTubesSuccess: "GET_TIPS_AND_TUBES_SUCCESS",
  getTipsAndTubesFailure: "GET_TIPS_AND_TUBES_FAILED",
  getTipsAndTubesReset: "GET_TIPS_AND_TUBES_RESET",
};

export const getCartridgeAction = {
  getCartridgeInitiated: "GET_CARTRIDGE_INITIATED",
  getCartridgeSuccess: "GET_CARTRIDGE_SUCCESS",
  getCartridgeFailure: "GET_CARTRIDGE_FAILED",
  getCartridgeReset: "GET_CARTRIDGE_RESET",
};

export default {
  updateRecipe: "UPDATE_RECIPE",
  saveRecipe: "SAVE_RECIPE",
  resetRecipe: "RESET_RECIPE",
};
