export const saveNewRecipeAction = {
  saveRecipeInitiated: "SAVE_RECIPE_INITIATE",
  saveRecipeName: "SAVE_RECIPE_NAME",
  saveRecipeSuccess: "SAVE_RECIPE_SUCCESS",
  saveRecipeFailed: "SAVE_RECIPE_FAILED",
  saveRecipeReset: "SAVE_RECIPE_RESET",
};

export const getTipsAndTubesAction = {
  getTipsAndTubesInitiated: "GET_TIPS_AND_TUBES_INITIATED",
  getTipsAndTubesSuccess: "GET_TIPS_AND_TUBES_SUCCESS",
  getTipsAndTubesFailure: "GET_TIPS_AND_TUBES_FAILED",
  getTipsAndTubesReset: "GET_TIPS_AND_TUBES_RESET",
}

export const getCartridgeAction = {
  getCartridgeInitiated: "GET_CARTRIDGE_INITIATED",
  getCartridgeSuccess: "GET_CARTRIDGE_SUCCESS",
  getCartridgeFailure: "GET_CARTRIDGE_FAILED",
  getCartridgeReset: "GET_CARTRIDGE_RESET",
}

export default {
  updateRecipe: "UPDATE_RECIPE",
  saveRecipe: "SAVE_RECIPE",
  resetRecipe: "RESET_RECIPE",
};
