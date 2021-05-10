import {
  saveNewRecipeAction,
  getTipsAndTubesAction,
  getCartridgeAction,
} from "actions/saveNewRecipeActions";

export const saveNewRecipe = (params) => ({
  type: saveNewRecipeAction.saveRecipeName,
  payload: params,
});

export const getTipsAndTubesActionInitiated = (params) => ({
  type: getTipsAndTubesAction.getTipsAndTubesInitiated,
  payload: params
});

export const getTipsAndTubesActionSuccess = (params) => ({
  type: getTipsAndTubesAction.getTipsAndTubesSuccess,
  payload: params
});

export const getTipsAndTubesActionFailure = (params) => ({
  type: getTipsAndTubesAction.getTipsAndTubesFailure,
  payload: params
});

export const getTipsAndTubesActionReset = (params) => ({
  type: getTipsAndTubesAction.getTipsAndTubesReset,
  payload: params
});

export const getCartridgeActionInitiated = (params) => ({
  type: getCartridgeAction.getCartridgeInitiated,
  payload: params
});

export const getCartridgeActionSuccess = (params) => ({
  type: getCartridgeAction.getCartridgeSuccess,
  payload: params
});

export const getCartridgeActionFailure = (params) => ({
  type: getCartridgeAction.getCartridgeActionFailure,
  payload: params
});

export const getCartridgeActionReset = (params) => ({
  type: getCartridgeAction.getCartridgeActionSuccess,
  payload: params
});

// //create/edit recipe detials
// export const updateRecipe = (payload) => ({
//     type: saveNewRecipeActions.updateRecipe,
//     payload
// });

// //clear all data
// export const resetRecipe = () => ({
//     type: saveNewRecipeActions.resetRecipe
// })
