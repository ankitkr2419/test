import {
  saveNewRecipeAction,
  getTipsAndTubesAction,
  getCartridgeAction,
  getTubesAction,
  getTipsAction,
} from "actions/saveNewRecipeActions";

export const saveNewRecipe = (params) => ({
  type: saveNewRecipeAction.saveRecipeName,
  payload: params,
});

//posting new recipe
export const updateRecipeActionInitiated = (params) => ({
  type: saveNewRecipeAction.updateRecipeInitiated,
  payload: params,
});

export const updateRecipeActionSuccess = (params) => ({
  type: saveNewRecipeAction.updateRecipeSuccess,
  payload: params,
});

export const updateRecipeActionFailure = (params) => ({
  type: saveNewRecipeAction.updateRecipeFailure,
  payload: params,
});

export const updateRecipeActionReset = (params) => ({
  type: saveNewRecipeAction.updateRecipeReset,
  payload: params,
});

// getting and saving tips and tubes options
export const getTipsAndTubesActionInitiated = (params) => ({
  type: getTipsAndTubesAction.getTipsAndTubesInitiated,
  payload: params,
});

export const getTipsAndTubesActionSuccess = (params) => ({
  type: getTipsAndTubesAction.getTipsAndTubesSuccess,
  payload: params,
});

export const getTipsAndTubesActionFailure = (params) => ({
  type: getTipsAndTubesAction.getTipsAndTubesFailure,
  payload: params,
});

export const getTipsAndTubesActionReset = (params) => ({
  type: getTipsAndTubesAction.getTipsAndTubesReset,
  payload: params,
});

// getting and saving cartridge options
export const getCartridgeActionInitiated = (params) => ({
  type: getCartridgeAction.getCartridgeInitiated,
  payload: params,
});

export const getCartridgeActionSuccess = (params) => ({
  type: getCartridgeAction.getCartridgeSuccess,
  payload: params,
});

export const getCartridgeActionFailure = (params) => ({
  type: getCartridgeAction.getCartridgeActionFailure,
  payload: params,
});

export const getCartridgeActionReset = (params) => ({
  type: getCartridgeAction.getCartridgeActionSuccess,
  payload: params,
});

// getting and saving tubes options
export const getTubesActionInitiated = (params) => ({
  type: getTubesAction.getTubesInitiated,
  payload: params,
});

export const getTubesActionSuccess = (params) => ({
  type: getTubesAction.getTubesSuccess,
  payload: params,
});

export const getTubesActionFailure = (params) => ({
  type: getTubesAction.getTubesFailure,
  payload: params,
});

export const getTubesActionReset = (params) => ({
  type: getTubesAction.getTubesReset,
  payload: params,
});

// getting and saving tips options
export const getTipsActionInitiated = (params) => ({
  type: getTipsAction.getTipsInitiated,
  payload: params,
});

export const getTipsActionSuccess = (params) => ({
  type: getTipsAction.getTipsSuccess,
  payload: params,
});

export const getTipsActionFailure = (params) => ({
  type: getTipsAction.getTipsFailure,
  payload: params,
});

export const getTipsActionReset = (params) => ({
  type: getTipsAction.getTipsReset,
  payload: params,
});
