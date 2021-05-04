import saveNewRecipeActions from "actions/saveNewRecipeActions";
import { fromJS } from "immutable";

const saveRecipeInitialState = fromJS({
    name: "",
    //isSaved
    //deckName
    //others...
});

export const saveNewRecipeReducer = (
    state = saveRecipeInitialState,
    action
) => {
    switch (action.type) {
        case saveNewRecipeActions.updateRecipe:
            return state.merge({
                ...action.payload
            });
        case saveNewRecipeActions.resetRecipe: 
            return saveRecipeInitialState;
        default:
            return state;
    }
};
