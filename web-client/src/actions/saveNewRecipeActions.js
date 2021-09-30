export const storeExistingRecipeAction = {
  saveRecipeDetails: "SAVE_RECIPE_DETAILS",
};

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

export const getTipsAction = {
  getTipsInitiated: "GET_TIPS_OPTIONS_INITIATED",
  getTipsSuccess: "GET_TIPS_OPTIONS_SUCCESS",
  getTipsFailure: "GET_TIPS_OPTIONS_FAILED",
  getTipsReset: "GET_TIPS_OPTIONS_RESET",
};

export const getTubesAction = {
  getTubesInitiated: "GET_TUBES_OPTIONS_INITIATED",
  getTubesSuccess: "GET_TUBES_OPTIONS_SUCCESS",
  getTubesFailure: "GET_TUBES_OPTIONS_FAILED",
  getTubesReset: "GET_TUBES_OPTIONS_RESET",
};

export const getCartridgeAction = {
  getCartridgeInitiated: "GET_CARTRIDGE_INITIATED",
  getCartridgeSuccess: "GET_CARTRIDGE_SUCCESS",
  getCartridgeFailure: "GET_CARTRIDGE_FAILED",
  getCartridgeReset: "GET_CARTRIDGE_RESET",
};

export const getCartridge1Action = {
  getCartridge1Initiated: "GET_CARTRIDGE_1_INITIATED",
  getCartridge1Success: "GET_CARTRIDGE_1_SUCCESS",
  getCartridge1Failure: "GET_CARTRIDGE_1_FAILED",
  getCartridge1Reset: "GET_CARTRIDGE_1_RESET",
};

export const getCartridge2Action = {
  getCartridge2Initiated: "GET_CARTRIDGE_2_INITIATED",
  getCartridge2Success: "GET_CARTRIDGE_2_SUCCESS",
  getCartridge2Failure: "GET_CARTRIDGE_2_FAILED",
  getCartridge2Reset: "GET_CARTRIDGE_2_RESET",
};

export default {
  updateRecipe: "UPDATE_RECIPE",
  saveRecipe: "SAVE_RECIPE",
  resetRecipe: "RESET_RECIPE",
};
