import {
  saveNewRecipeAction,
  getRecipeDetailsAction,
} from "actions/saveNewRecipeActions";

export const saveNewRecipe = (params) => ({
  type: saveNewRecipeAction.saveRecipeName,
  payload: params,
});

export const getRecipeDetailsActionInitiaed = (params) => ({
  type: getRecipeDetailsAction.getRecipeDetailsInitiated,
  payload: params
});

export const getRecipeDetailsActionSuccess = (params) => ({
  type: getRecipeDetailsAction.getRecipeDetailsSuccess,
  payload: params
});

export const getRecipeDetailsActionFailure = (params) => ({
  type: getRecipeDetailsAction.getRecipeDetailsFailure,
  payload: params
});

export const getRecipeDetailsActionReset = (params) => ({
  type: getRecipeDetailsAction.getRecipeDetailsReset,
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
