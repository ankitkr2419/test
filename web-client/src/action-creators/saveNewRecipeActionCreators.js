import saveNewRecipeActions from "actions/saveNewRecipeActions";

//create/edit recipe detials
export const updateRecipe = (payload) => ({
    type: saveNewRecipeActions.updateRecipe,
    payload
});

//clear all data
export const resetRecipe = () => ({
    type: saveNewRecipeActions.resetRecipe
})
