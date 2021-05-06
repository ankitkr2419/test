import { saveNewRecipeAction } from "actions/saveNewRecipeActions";

export const saveNewRecipe = (params) => ({
  type: saveNewRecipeAction.saveRecipeName,
  payload: params,
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
